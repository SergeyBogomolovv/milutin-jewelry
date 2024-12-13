package controller

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/SergeyBogomolovv/milutin-jewelry/internal/domain/dto"
	errs "github.com/SergeyBogomolovv/milutin-jewelry/internal/domain/errors"
	"github.com/SergeyBogomolovv/milutin-jewelry/internal/middleware"
	"github.com/SergeyBogomolovv/milutin-jewelry/pkg/utils"
	"github.com/go-playground/validator/v10"
)

type CollectionsUsecase interface {
	CreateCollection(context.Context, *dto.CreateCollectionRequest) (int, error)
	UpdateCollection(context.Context, *dto.UpdateCollectionRequest) error
	GetAllCollections(context.Context) ([]*dto.CollectionResponse, error)
	GetCollectionByID(context.Context, int) (*dto.CollectionResponse, error)
}

type collectionsController struct {
	validate *validator.Validate
	uc       CollectionsUsecase
	log      *slog.Logger
}

func RegisterCollectionsController(log *slog.Logger, router *http.ServeMux, uc CollectionsUsecase, authMw middleware.Middleware) {
	controller := &collectionsController{
		log:      log.With(slog.String("op", "collectionsController")),
		uc:       uc,
		validate: validator.New(validator.WithRequiredStructEnabled()),
	}

	router.HandleFunc("GET /collections/all", controller.GetAllCollections)
	router.HandleFunc("GET /collections/{id}", controller.GetOneCollection)

	admin := http.NewServeMux()
	admin.HandleFunc("POST /create", controller.CreateCollection)
	admin.HandleFunc("PUT /update/{id}", controller.UpdateCollection)
	admin.HandleFunc("DELETE /delete/{id}", controller.DeleteCollection)
	router.Handle("/collections/", http.StripPrefix("/collections", authMw(admin)))
}

func (c *collectionsController) CreateCollection(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		c.log.Error("failed to parse multipart form", "err", err)
		utils.WriteError(w, "invalid form", http.StatusBadRequest)
		return
	}
	file, _, err := r.FormFile("image")
	if err != nil {
		c.log.Error("failed to get image", "err", err)
		utils.WriteError(w, "no image provided", http.StatusBadRequest)
		return
	}
	defer file.Close()

	dto := &dto.CreateCollectionRequest{Title: r.FormValue("title"), Description: r.FormValue("description"), Image: file}
	if err := c.validate.Struct(dto); err != nil {
		c.log.Error("failed to validate payload", "err", err)
		utils.WriteError(w, "invalid payload", http.StatusBadRequest)
		return
	}
	c.log.Info("creating collection", "title", dto.Title)

	id, err := c.uc.CreateCollection(r.Context(), dto)
	if err != nil {
		c.log.Error("failed to create collection", "err", err)
		utils.WriteError(w, "failed to create collection", http.StatusInternalServerError)
		return
	}
	utils.WriteJSON(w, map[string]int{"collection_id": id}, http.StatusCreated)
}

func (c *collectionsController) UpdateCollection(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		c.log.Error("failed to parse multipart form", "err", err)
		utils.WriteError(w, "invalid form", http.StatusBadRequest)
		return
	}
	dto := &dto.UpdateCollectionRequest{
		Title:       r.FormValue("title"),
		Description: r.FormValue("description"),
	}
	if file, _, err := r.FormFile("image"); err == nil {
		dto.Image = file
		defer file.Close()
	}

	if err := c.validate.Struct(dto); err != nil {
		c.log.Error("failed to validate payload", "err", err)
		utils.WriteError(w, "invalid payload", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		c.log.Error("failed to parse id", "err", err)
		utils.WriteError(w, "invalid id", http.StatusBadRequest)
		return
	}
	dto.ID = id

	if err := c.uc.UpdateCollection(r.Context(), dto); err != nil {
		if errors.Is(err, errs.ErrCollectionNotFound) {
			c.log.Info("collection not found", "err", err)
			utils.WriteError(w, "collection not found", http.StatusNotFound)
			return
		}
		c.log.Error("failed to update collection", "err", err)
		utils.WriteError(w, "failed to update collection", http.StatusInternalServerError)
		return
	}

	utils.WriteMessage(w, "collection updated", http.StatusCreated)
}

func (c *collectionsController) DeleteCollection(w http.ResponseWriter, r *http.Request) {}

func (c *collectionsController) GetAllCollections(w http.ResponseWriter, r *http.Request) {
	collections, err := c.uc.GetAllCollections(r.Context())
	if err != nil {
		c.log.Error("failed to get collections", "err", err)
		utils.WriteError(w, "failed to get collections", http.StatusInternalServerError)
		return
	}
	utils.WriteJSON(w, collections, http.StatusOK)
}

func (c *collectionsController) GetOneCollection(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		c.log.Error("failed to parse id", "err", err)
		utils.WriteError(w, "invalid id", http.StatusBadRequest)
		return
	}
	collection, err := c.uc.GetCollectionByID(r.Context(), id)
	if err != nil {
		if errors.Is(err, errs.ErrCollectionNotFound) {
			c.log.Info("collection not found", "err", err)
			utils.WriteError(w, "collection not found", http.StatusNotFound)
			return
		}
		c.log.Error("failed to get collection", "err", err)
		utils.WriteError(w, "failed to get collection", http.StatusInternalServerError)
		return
	}
	utils.WriteJSON(w, collection, http.StatusOK)
}
