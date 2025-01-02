package app

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/SergeyBogomolovv/milutin-jewelry/internal/config"
	"github.com/SergeyBogomolovv/milutin-jewelry/internal/controller"
	authcontroller "github.com/SergeyBogomolovv/milutin-jewelry/internal/controller/auth"
	itemscontroller "github.com/SergeyBogomolovv/milutin-jewelry/internal/controller/items"
	"github.com/SergeyBogomolovv/milutin-jewelry/internal/infra"
	"github.com/SergeyBogomolovv/milutin-jewelry/internal/infra/files"
	"github.com/SergeyBogomolovv/milutin-jewelry/internal/middleware"
	repo "github.com/SergeyBogomolovv/milutin-jewelry/internal/storage"
	codestorage "github.com/SergeyBogomolovv/milutin-jewelry/internal/storage/codes"
	itemstorage "github.com/SergeyBogomolovv/milutin-jewelry/internal/storage/items"
	"github.com/SergeyBogomolovv/milutin-jewelry/internal/usecase"
	authusecase "github.com/SergeyBogomolovv/milutin-jewelry/internal/usecase/auth"
	itemsusecase "github.com/SergeyBogomolovv/milutin-jewelry/internal/usecase/items"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

type application struct {
	srv *http.Server
	log *slog.Logger
}

func New(log *slog.Logger, db *sqlx.DB, redis *redis.Client, cfg *config.Config) *application {
	router := http.NewServeMux()
	router.Handle("/docs/", httpSwagger.WrapHandler)
	authMiddleware := middleware.NewAuthMiddleware(cfg.JWTSecret)
	corsMiddleware := middleware.NewCORSMiddleware(cfg.CORSOrigin)
	loggerMiddleware := middleware.NewLoggerMiddleware(log)
	filesService := files.New(log, cfg.ObjectStorage)

	collectionsRepo := repo.NewCollectionsRepo(db)
	collectionsUsecase := usecase.NewCollectionsUsecase(log, filesService, collectionsRepo)
	controller.RegisterCollectionsController(log, router, collectionsUsecase, authMiddleware)

	itemsStorage := itemstorage.New(db)
	itemsUsecase := itemsusecase.New(log, filesService, itemsStorage)
	itemscontroller.Register(log, router, itemsUsecase, authMiddleware)

	mailService := infra.NewMailService(log, cfg.Mail, cfg.AdminEmail)
	codeStorage := codestorage.New(redis)
	authUsecase := authusecase.New(log, codeStorage, mailService, cfg.JWTSecret)
	authcontroller.Register(log, router, authUsecase)

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
