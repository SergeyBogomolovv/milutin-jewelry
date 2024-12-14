package app

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/SergeyBogomolovv/milutin-jewelry/internal/config"
	"github.com/SergeyBogomolovv/milutin-jewelry/internal/controller"
	"github.com/SergeyBogomolovv/milutin-jewelry/internal/infra"
	"github.com/SergeyBogomolovv/milutin-jewelry/internal/middleware"
	"github.com/SergeyBogomolovv/milutin-jewelry/internal/repo"
	"github.com/SergeyBogomolovv/milutin-jewelry/internal/usecase"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

type application struct {
	srv *http.Server
	log *slog.Logger
}

func New(log *slog.Logger, db *sqlx.DB, redis *redis.Client, cfg *config.Config) *application {
	router := http.NewServeMux()
	authMiddleware := middleware.NewAuthMiddleware(cfg.JWTSecret)
	filesService := infra.NewFilesService(log, cfg.ObjectStorage)

	collectionsRepo := repo.NewCollectionsRepo(db)
	collectionsUsecase := usecase.NewCollectionsUsecase(log, filesService, collectionsRepo)
	controller.RegisterCollectionsController(log, router, collectionsUsecase, authMiddleware)

	mailService := infra.NewMailService(log, cfg.Mail, cfg.AdminEmail)
	codesRepo := repo.NewCodesRepo(redis)
	authUsecase := usecase.NewAuthUsecase(log, codesRepo, mailService, cfg.JWTSecret)
	controller.RegisterAuthController(log, router, authUsecase)

	return &application{
		srv: &http.Server{Addr: cfg.Addr, Handler: router},
		log: log,
	}
}

func (a *application) Start() {
	const op = "app.Run"
	log := a.log.With(slog.String("op", op))
	log.Info("starting server", "addr", a.srv.Addr)
	a.srv.ListenAndServe()
}

func (a *application) Stop() {
	const op = "app.Stop"
	log := a.log.With(slog.String("op", op))
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	a.srv.Shutdown(ctx)
	log.Info("server stopped")
}
