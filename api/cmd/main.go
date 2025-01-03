package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/SergeyBogomolovv/milutin-jewelry/docs"
	"github.com/SergeyBogomolovv/milutin-jewelry/internal/app"
	"github.com/SergeyBogomolovv/milutin-jewelry/internal/config"
	"github.com/SergeyBogomolovv/milutin-jewelry/pkg/db"
	"github.com/SergeyBogomolovv/milutin-jewelry/pkg/redis"
	"github.com/joho/godotenv"
)

// @title Milutin Jewelry API
// @version 1.0
// @description Описание API для сервиса Milutin Jewelry
func main() {
	cfg := config.New()

	db := db.New(cfg.PostgresUrl)
	defer db.Close()
	redis := redis.New(cfg.RedisUrl)
	defer redis.Close()

	log := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	app := app.New(log, db, redis, cfg)

	go app.Start()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	<-ctx.Done()
	app.Stop()
}

func init() {
	godotenv.Load()
}
