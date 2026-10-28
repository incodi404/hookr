package db

import (
	"context"
	"fmt"
	"time"

	logruslogger "github.com/incodi404/hookr.git/pkg/logrus-logger"
	"github.com/incodi404/hookr.git/pkg/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

var DbPool *pgxpool.Pool

func ConnectPGPool(ctx context.Context, config models.DBConfig) (*pgxpool.Pool, error) {
	cfg, err := pgxpool.ParseConfig("")
	if err != nil {
		logruslogger.Writer.Error("[PGSQL] Error occured in PGSQL: ", err)
		return nil, err
	}

	cfg.ConnConfig.Host = config.Host
	cfg.ConnConfig.Port = uint16(config.Port)
	cfg.ConnConfig.Database = config.DbName
	cfg.ConnConfig.User = config.User
	cfg.ConnConfig.Password = config.Password

	// Pool tuning
	cfg.MaxConns = 20
	cfg.MinConns = 5
	cfg.MaxConnLifetime = 30 * time.Minute
	cfg.MaxConnIdleTime = 5 * time.Minute
	cfg.HealthCheckPeriod = time.Minute

	pool, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		logruslogger.Writer.Error("[PGSQL] Error occured in PGSQL: ", err)
		return nil, err
	}

	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		logruslogger.Writer.Error("[PGSQL] Error occured in PGSQL: ", err)
		return nil, err
	}

	DbPool = pool

	logruslogger.Writer.Info("[PGSQL] PGSQL Database is connected")
	fmt.Println("[PGSQL] PGSQL Database is connected")

	return pool, nil
}
