package controller

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/SergeyBogomolovv/milutin-jewelry/internal/domain/dto"
	"github.com/SergeyBogomolovv/milutin-jewelry/internal/middleware"
	"github.com/go-playground/validator/v10"
)

type CollectionItemsUsecase interface {
	Create(context.Context, *dto.CreateCollectionItemRequest) (int, error)
	Update(context.Context, *dto.UpdateCollectionItemRequest) error
	Delete(context.Context, int) error
	GetByCollection(context.Context, int) ([]*dto.CollectionItemResponse, error)
	GetOne(context.Context, int) (*dto.CollectionItemResponse, error)
}

type collectionItemsController struct {
	uc       CollectionItemsUsecase
	log      *slog.Logger
	validate *validator.Validate
}

func RegisterCollectionItemsController(log *slog.Logger, router *http.ServeMux, uc CollectionItemsUsecase, auth middleware.Middleware) {
	c := &collectionItemsController{
		log:      log.With(slog.String("op", "collectionItemsController")),
		uc:       uc,
		validate: validator.New(validator.WithRequiredStructEnabled()),
	}
	r := http.NewServeMux()
	r.HandleFunc("GET /one/{id}", c.GetOneCollectionItem)
	r.HandleFunc("GET /collection/{id}", c.GetCollectionItemsByCollection)

	pr := http.NewServeMux()
	pr.HandleFunc("POST /create", c.CreateCollectionItem)
	pr.HandleFunc("PUT /update/{id}", c.UpdateCollectionItem)
	pr.HandleFunc("DELETE /delete/{id}", c.DeleteCollectionItem)
	r.Handle("/", auth(pr))

	router.Handle("/collection-items/", http.StripPrefix("/collection-items", r))
}

func (c *collectionItemsController) CreateCollectionItem(w http.ResponseWriter, r *http.Request) {
}

func (c *collectionItemsController) UpdateCollectionItem(w http.ResponseWriter, r *http.Request) {}

func (c *collectionItemsController) DeleteCollectionItem(w http.ResponseWriter, r *http.Request) {}

func (c *collectionItemsController) GetCollectionItemsByCollection(w http.ResponseWriter, r *http.Request) {
}

func (c *collectionItemsController) GetOneCollectionItem(w http.ResponseWriter, r *http.Request) {}
