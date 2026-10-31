package redis

import (
	"context"
	"fmt"

	logruslogger "github.com/incodi404/hookr.git/pkg/logrus-logger"
	"github.com/incodi404/hookr.git/pkg/models"
	"github.com/redis/go-redis/v9"
)

var RdbCli *redis.Client

type RedisClient struct {
	Client *redis.Client
}

func ConnectRedis(ctx context.Context, cfg models.RdbConfig) (*RedisClient, error) {
	// create connection
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Host + ":" + cfg.Port,
		Password: cfg.Password,
		DB:       cfg.Db,
	})

	err := rdb.Ping(ctx).Err()
	if err != nil {
		logruslogger.Writer.Error("[REDIS] Redis connection has been failed: ", err)
		return nil, err
	}

	logruslogger.Writer.Info("[REDIS] Redis is connected")
	fmt.Println("[REDIS] Redis is connected")

	return &RedisClient{
		Client: rdb,
	}, nil
}

func (r *RedisClient) Close() error {
	if r == nil || r.Client == nil {
		return nil
	}

	return r.Client.Close()
}

func (r *RedisClient) ValidateClient() bool {
	if r == nil || r.Client == nil {
		return false
	}

	return true
}
