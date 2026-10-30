package redis

import (
	"context"
	"fmt"
	"time"
)

func SetRdbHash(ctx context.Context, key string, data interface{}, exp *time.Duration) error {
	if RdbCli == nil {
		return fmt.Errorf("[REDIS] Redis is not connected")
	}

	// set the hash
	err := RdbCli.HSet(ctx, key, data).Err()
	if err != nil {
		return err
	}

	// set expiration (optional)
	if exp != nil {
		err = RdbCli.Expire(ctx, key, *exp).Err()
		if err != nil {
			return err
		}
	}

	return nil
}
