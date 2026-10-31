package repositories

import (
	"context"
	"time"

	"github.com/incodi404/hookr.git/pkg/redis"
)

type OTPRepository interface {
	StoreOtp(
		ctx context.Context,
		email_id string,
		hashed_otp string,
	) error

	GetOtp(
		ctx context.Context,
		email_id string,
	) (string, error)
}

type RedisOTPRepository struct {
	Client *redis.RedisClient
}

func NewRedisOTPRepository(
	client *redis.RedisClient,
) *RedisOTPRepository {
	return &RedisOTPRepository{
		Client: client,
	}
}

func (r *RedisOTPRepository) StoreOtp(
	ctx context.Context,
	email_id string,
	hashed_otp string,
) error {

	ttl := 10 * time.Minute // TTL for saving otp
	err := r.Client.SetHash(ctx, "verify:"+email_id, map[string]interface{}{
		"hashed_otp": hashed_otp,
	}, &ttl)

	if err != nil {
		return err
	}

	return nil
}

func (r *RedisOTPRepository) GetOtp(
	ctx context.Context,
	email_id string,
) (string, error) {

	hashed_otp, err := r.Client.GetHash(ctx, "verify:"+email_id, "hashed_otp")
	if err != nil {
		return "", nil
	}

	return hashed_otp, nil
}
