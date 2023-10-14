package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
)

var redisClient *redis.Client

func InitRedisCache(ctx context.Context) error {
	client, err := InitClient(ctx)
	if err != nil {
		// log
		return err
	}
	redisClient = client
	return nil
}
func InitClient(ctx context.Context) (*redis.Client, error) {
	option := &redis.Options{
		Addr:     "localhost",
		DB:       0,
		PoolSize: 20,
		Password: "",
	}
	client := redis.NewClient(option)
	_, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}
	return client, nil
}
