package main

import (
	"context"
	"log"
	"time"

	api "github.com/incodi404/hookr.git/internal/api/src"
	"github.com/incodi404/hookr.git/pkg/db"
	logruslogger "github.com/incodi404/hookr.git/pkg/logrus-logger"
	"github.com/incodi404/hookr.git/pkg/models"
	"github.com/incodi404/hookr.git/pkg/redis"
	"github.com/incodi404/hookr.git/pkg/utils"
	"github.com/joho/godotenv"
)

func main() {
	// config logrus
	logruslogger.ConfigLogrus()

	err := godotenv.Load()
	if err != nil {
		logruslogger.Writer.Error("[ENV] Error loading .env file")
		log.Fatal("Error loading .env file")
	}

	ctx := context.Background()

	// connecting pg db
	_, _, err = db.ConnectPGPool(ctx, models.DBConfig{
		Host:     utils.GetEnv("POSTGRES_HOST", "hookr_db"),
		Port:     utils.StrToInt(utils.GetEnv("POSTGRES_PORT", "5432")),
		DbName:   utils.GetEnv("POSTGRES_DB", "hookr_db"),
		User:     utils.GetEnv("POSTGRES_USER", "NONE"),
		Password: utils.GetEnv("POSTGRES_PASSWORD", "NONE"),
	})
	if err != nil {
		logruslogger.Writer.Error(err)
		log.Fatal(err)
	}

	// connecting redis
	_, err = redis.ConnectRedis(ctx, models.RdbConfig{
		Host:     utils.GetEnv("REDIS_HOST", "redis"),
		Port:     utils.GetEnv("REDIS_PORT", "6379"),
		Password: utils.GetEnv("REDIS_PASSWORD", "NONE"),
		Db:       0,
	})
	if err != nil {
		logruslogger.Writer.Error("[REDIS] Redis connection has been failed: ", err)
		log.Fatal("[REDIS] Redis connection has been failed: ", err)
	}

	// starting server
	_, err = api.StartServer(models.ServerConfig{
		Addr:           ":7000",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	})

	if err != nil {
		logruslogger.Writer.Error(err)
		log.Fatal(err)
	}
}
