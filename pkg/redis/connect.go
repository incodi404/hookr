package redis

import (
	"context"

	logruslogger "github.com/incodi404/hookr.git/pkg/logrus-logger"
	"github.com/incodi404/hookr.git/pkg/models"
	"github.com/redis/go-redis/v9"
)

var RdbCli *redis.Client

func ConnectRedis(ctx context.Context, cfg models.RdbConfig) (*redis.Client, error) {
	// create connection
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Host + cfg.Port,
		Password: cfg.Password,
		DB:       cfg.Db,
	})

	err := rdb.Ping(ctx).Err()
	if err != nil {
		logruslogger.Writer.Error(err)
		return nil, err
	}

	RdbCli = rdb

	return rdb, nil
}
