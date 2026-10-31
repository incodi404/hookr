package redis

import (
	"context"
	"fmt"
	"time"
)

func (r *RedisClient) SetHash(ctx context.Context, key string, data interface{}, exp *time.Duration) error {
	if !r.ValidateClient() {
		return fmt.Errorf("[REDIS] Redis is not connected")
	}

	// set the hash
	err := r.Client.HSet(ctx, key, data).Err()
	if err != nil {
		return err
	}

	// set expiration (optional)
	if exp != nil {
		err = r.Client.Expire(ctx, key, *exp).Err()
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *RedisClient) GetHash(
	ctx context.Context,
	key string,
	field string,
) (string, error) {
	if !r.ValidateClient() {
		return "", fmt.Errorf("[REDIS] Redis is not connected")
	}

	data, err := r.Client.HGet(ctx, key, field).Result()
	if err != nil {
		return "", nil
	}

	return data, nil
}

func (r *RedisClient) GetHashAll(
	ctx context.Context,
	key string,
) (map[string]string, error) {
	data := make(map[string]string)

	if !r.ValidateClient() {
		return data, fmt.Errorf("[REDIS] Redis is not connected")
	}

	data, err := r.Client.HGetAll(ctx, key).Result()
	if err != nil {
		return data, err
	}

	return data, nil
}
