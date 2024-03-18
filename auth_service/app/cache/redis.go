package cache

import (
	"authservice/app/config"
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

func InitRedis() Redis {
	opt, err := redis.ParseURL(config.RDS_URL)
	if err != nil {
		panic(err)
	}
	rdb := redis.NewClient(opt)
	return &redisClient{
		rdb: rdb,
	}
}

type redisClient struct {
	rdb *redis.Client
}

type Redis interface {
	Set(ctx context.Context, key string, value string) error
	Get(ctx context.Context, key string) (string, error)
	Delete(ctx context.Context, key string) error
}

func (c *redisClient) Set(ctx context.Context, key string, value string) error {
	err := c.rdb.Set(ctx, key, value, 10*time.Minute).Err()
	return err
}

func (c *redisClient) Get(ctx context.Context, key string) (string, error) {
	val, err := c.rdb.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}

	return val, nil
}

func (c *redisClient) Delete(ctx context.Context, key string) error {
	err := c.rdb.Del(ctx, key).Err()
	return err
}
