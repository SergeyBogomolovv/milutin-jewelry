package redis

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

func New(url string) *redis.Client {
	opts, err := redis.ParseURL(url)
	if err != nil {
		log.Fatalf("failed to parse redis url: %s", err)
	}
	client := redis.NewClient(opts)

	ping, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("failed to connect to redis: %s", err)
	}
	if ping != "PONG" {
		log.Fatalf("failed to ping redis: %s", ping)
	}

	return client
}
