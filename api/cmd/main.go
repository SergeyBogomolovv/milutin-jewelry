package main

import (
	"github.com/SergeyBogomolovv/milutin-jewelry/internal/config"
	"github.com/SergeyBogomolovv/milutin-jewelry/pkg/db"
	"github.com/SergeyBogomolovv/milutin-jewelry/pkg/redis"
)

func main() {
	cfg := config.New()

	db := db.New(cfg.PostgresUrl)
	defer db.Close()
	redis := redis.New(cfg.RedisUrl)
	defer redis.Close()
}
