package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"sync"
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
	redis := redis.New(cfg.RedisUrl)

	log := newLogger()

	app := app.New(log, db, redis, cfg)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	defer stop()

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		<-ctx.Done()
		app.Stop()
	}()

	app.Start()
	wg.Wait()
}

func init() {
	godotenv.Load()
}

func newLogger() *slog.Logger {
	return slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
}
