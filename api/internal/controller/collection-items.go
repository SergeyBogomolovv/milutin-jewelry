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
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		c.log.Error("failed to parse multipart form", "err", err)
		utils.WriteError(w, "invalid form", http.StatusBadRequest)
		return
	}
	collectionID, err := strconv.Atoi(r.FormValue("collection_id"))
	if err != nil {
		utils.WriteError(w, "invalid collection id", http.StatusBadRequest)
		return
	}
	file, _, err := r.FormFile("image")
	if err != nil {
		utils.WriteError(w, "no image provided", http.StatusBadRequest)
		return
	}
	defer file.Close()
	dto := &dto.CreateCollectionItemRequest{
		CollectionID: collectionID,
		Title:        r.FormValue("title"),
		Description:  r.FormValue("description"),
		Image:        file,
	}
	if err := c.validate.Struct(dto); err != nil {
		utils.WriteError(w, "invalid payload", http.StatusBadRequest)
		return
	}
	id, err := c.uc.Create(r.Context(), dto)
	if err != nil {
		if errors.Is(err, errs.ErrCollectionNotFound) {
			utils.WriteError(w, "collection not found", http.StatusNotFound)
			return
		}
		utils.WriteError(w, "failed to create collection item", http.StatusInternalServerError)
		return
	}
	utils.WriteJSON(w, map[string]int{"collection_item_id": id}, http.StatusCreated)
}

func (c *collectionItemsController) UpdateCollectionItem(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		c.log.Error("failed to parse multipart form", "err", err)
		utils.WriteError(w, "invalid form", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		utils.WriteError(w, "invalid id", http.StatusBadRequest)
		return
	}
	dto := &dto.UpdateCollectionItemRequest{
		ID:          id,
		Title:       r.FormValue("title"),
		Description: r.FormValue("description"),
	}
	if file, _, err := r.FormFile("image"); err == nil {
		dto.Image = file
		defer file.Close()
	}

	if err := c.validate.Struct(dto); err != nil {
		utils.WriteError(w, "invalid payload", http.StatusBadRequest)
		return
	}

	if err := c.uc.Update(r.Context(), dto); err != nil {
		if errors.Is(err, errs.ErrCollectionItemNotFound) {
			utils.WriteError(w, "collection item not found", http.StatusNotFound)
			return
		}
		utils.WriteError(w, "failed to update collection item", http.StatusInternalServerError)
		return
	}
	utils.WriteMessage(w, "collection-item updated", http.StatusCreated)
}

func (c *collectionItemsController) DeleteCollectionItem(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		utils.WriteError(w, "invalid id", http.StatusBadRequest)
		return
	}
	if err := c.uc.Delete(r.Context(), id); err != nil {
		if errors.Is(err, errs.ErrCollectionItemNotFound) {
			utils.WriteError(w, "collection item not found", http.StatusNotFound)
			return
		}
		utils.WriteError(w, "failed to delete collection item", http.StatusInternalServerError)
		return
	}
	utils.WriteMessage(w, "collection-item deleted", http.StatusOK)
}

func (c *collectionItemsController) GetCollectionItemsByCollection(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		utils.WriteError(w, "invalid id", http.StatusBadRequest)
		return
	}
	collectionItems, err := c.uc.GetByCollection(r.Context(), id)
	if err != nil {
		if errors.Is(err, errs.ErrCollectionNotFound) {
			utils.WriteError(w, "collection not found", http.StatusNotFound)
			return
		}
		utils.WriteError(w, "failed to get collection items", http.StatusInternalServerError)
		return
	}
	utils.WriteJSON(w, collectionItems, http.StatusOK)
}

func (c *collectionItemsController) GetOneCollectionItem(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		utils.WriteError(w, "invalid id", http.StatusBadRequest)
		return
	}
	collectionItem, err := c.uc.GetOne(r.Context(), id)
	if err != nil {
		if errors.Is(err, errs.ErrCollectionItemNotFound) {
			utils.WriteError(w, "collection item not found", http.StatusNotFound)
			return
		}
		utils.WriteError(w, "failed to get collection item", http.StatusInternalServerError)
		return
	}
	utils.WriteJSON(w, collectionItem, http.StatusOK)
}
