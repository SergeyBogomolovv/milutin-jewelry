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
	repo "github.com/SergeyBogomolovv/milutin-jewelry/internal/storage"
	codestorage "github.com/SergeyBogomolovv/milutin-jewelry/internal/storage/codes"
	"github.com/SergeyBogomolovv/milutin-jewelry/internal/usecase"
	authusecase "github.com/SergeyBogomolovv/milutin-jewelry/internal/usecase/auth"
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
	corsMiddleware := middleware.NewCORSMiddleware(cfg.CORSOrigin)
	loggerMiddleware := middleware.NewLoggerMiddleware(log)
	filesService := infra.NewFilesService(log, cfg.ObjectStorage)

	collectionsRepo := repo.NewCollectionsRepo(db)
	collectionsUsecase := usecase.NewCollectionsUsecase(log, filesService, collectionsRepo)
	controller.RegisterCollectionsController(log, router, collectionsUsecase, authMiddleware)

	collectionItemsRepo := repo.NewCollectionItemsRepo(db)
	collectionItemsUsecase := usecase.NewCollectionItemsUsecase(log, filesService, collectionItemsRepo)
	controller.RegisterCollectionItemsController(log, router, collectionItemsUsecase, authMiddleware)

	mailService := infra.NewMailService(log, cfg.Mail, cfg.AdminEmail)
	codeStorage := codestorage.New(redis)
	authUsecase := authusecase.New(log, codeStorage, mailService, cfg.JWTSecret)
	controller.RegisterAuthController(log, router, authUsecase)

	return &application{
		srv: &http.Server{Addr: cfg.Addr, Handler: corsMiddleware(loggerMiddleware(router))},
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
