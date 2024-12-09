package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/SergeyBogomolovv/milutin-jewelry/internal/app"
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

	log := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	app := app.New(log, db, redis, cfg.Addr, cfg.JWTSecret)

	go app.Start()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	<-ctx.Done()
	app.Stop()
}
