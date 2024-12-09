package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

func New(url string) *redis.Client {
	opts, err := redis.ParseURL(url)
	if err != nil {
		panic(err)
	}
	client := redis.NewClient(opts)

	ping, err := client.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
	if ping != "PONG" {
		panic("failed to connect to redis")
	}

	return client
}
