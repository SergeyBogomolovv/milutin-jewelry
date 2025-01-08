package app

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/SergeyBogomolovv/milutin-jewelry/internal/config"
	authController "github.com/SergeyBogomolovv/milutin-jewelry/internal/controller/auth"
	collectionController "github.com/SergeyBogomolovv/milutin-jewelry/internal/controller/collection"
	itemController "github.com/SergeyBogomolovv/milutin-jewelry/internal/controller/item"
	fileService "github.com/SergeyBogomolovv/milutin-jewelry/internal/infra/file"
	"github.com/SergeyBogomolovv/milutin-jewelry/internal/infra/mail"
	"github.com/SergeyBogomolovv/milutin-jewelry/internal/middleware"
	codeStorage "github.com/SergeyBogomolovv/milutin-jewelry/internal/storage/code"
	collectionStorage "github.com/SergeyBogomolovv/milutin-jewelry/internal/storage/collection"
	itemStorage "github.com/SergeyBogomolovv/milutin-jewelry/internal/storage/item"
	authUsecase "github.com/SergeyBogomolovv/milutin-jewelry/internal/usecase/auth"
	collectionUsecase "github.com/SergeyBogomolovv/milutin-jewelry/internal/usecase/collection"
	itemUsecase "github.com/SergeyBogomolovv/milutin-jewelry/internal/usecase/item"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

type application struct {
	srv   *http.Server
	log   *slog.Logger
	db    *sqlx.DB
	redis *redis.Client
}

func New(log *slog.Logger, db *sqlx.DB, redis *redis.Client, cfg config.Config) *application {
	router := http.NewServeMux()
	router.Handle("/docs/", httpSwagger.WrapHandler)

	authMW := middleware.NewAuthMiddleware(cfg.Jwt.Secret)
	corsMW := middleware.NewCORSMiddleware(cfg.CORSOrigin)
	logMW := middleware.NewLoggerMiddleware(log)

	filesService := fileService.New(log, cfg.ObjectStorage)
	mailService := mail.New(log, cfg.Mail)

	collectionStorage := collectionStorage.New(db)
	collectionUsecase := collectionUsecase.New(log, filesService, collectionStorage)
	collectionController.Register(router, collectionUsecase, authMW)

	itemStorage := itemStorage.New(db)
	itemUsecase := itemUsecase.New(log, filesService, itemStorage)
	itemController.Register(router, itemUsecase, authMW)

	codeStorage := codeStorage.New(redis)
	authUsecase := authUsecase.New(log, codeStorage, mailService, cfg.Jwt)
	authController.Register(router, authUsecase)

	srv := &http.Server{Addr: cfg.Addr, Handler: corsMW(logMW(router))}

	return &application{
		srv:   srv,
		log:   log,
		db:    db,
		redis: redis,
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

	a.redis.Shutdown(ctx)
	log.Info("redis stopped")

	a.db.Close()
	log.Info("database stopped")
}
