package redis

//go:generate mockgen -destination=mocks/mock_redis_client.go -package=mocks . IRedisClient

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type IRedisClient interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
}

type client struct {
	redisClient *redis.Client
}

func NewRedisClient() IRedisClient {
	return &client{
		redisClient: redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
		}),
	}
}

func (c *client) Get(ctx context.Context, key string) (string, error) {
	return c.redisClient.Get(ctx, key).Result()
}

func (c *client) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return c.redisClient.Set(ctx, key, value, expiration).Err()
}
