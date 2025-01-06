package items

import (
	"errors"
	"log/slog"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/SergeyBogomolovv/milutin-jewelry/internal/middleware"
	uc "github.com/SergeyBogomolovv/milutin-jewelry/internal/usecase/item"
	"github.com/SergeyBogomolovv/milutin-jewelry/pkg/lib/res"
	"github.com/go-playground/validator/v10"
)

type controller struct {
	usecase  Usecase
	log      *slog.Logger
	validate *validator.Validate
}

func Register(log *slog.Logger, router *http.ServeMux, usecase Usecase, auth middleware.Middleware) {
	const dest = "itemsController"
	c := &controller{
		log:      log.With(slog.String("dest", dest)),
		usecase:  usecase,
		validate: validator.New(validator.WithRequiredStructEnabled()),
	}
	r := http.NewServeMux()
	r.HandleFunc("GET /collection/{id}", c.ItemsByCollection)
	r.HandleFunc("GET /{id}", c.ById)

	pr := http.NewServeMux()
	pr.HandleFunc("POST /create", c.CreateCollectionItem)
	pr.HandleFunc("PUT /update/{id}", c.UpdateCollectionItem)
	pr.HandleFunc("DELETE /delete/{id}", c.DeleteCollectionItem)
	r.Handle("/", auth(pr))

	router.Handle("/items/", http.StripPrefix("/items", r))
}

// @Summary      Создание украшения
// @Description  Создание украшения в коллекции
// @Tags         items
// @Accept       multipart/form-data
// @Produce      json
// @Param        collection_id  formData  int     true  "ID коллекции"
// @Param        title          formData  string  false  "Название"
// @Param        description    formData  string  false  "Описание"
// @Param        image          formData  file    true  "Изображение"
// @Success      201  {object}  Item
// @Failure      400  {object}  res.ErrorResponse  "Неверный формат данных"
// @Failure      401  {object}  res.ErrorResponse  "Нет доступа"
// @Failure      404  {object}  res.ErrorResponse  "Коллекции не существует"
// @Failure      500  {object}  res.ErrorResponse  "Ошибка сервера"
// @Router       /items/create [post]
func (c *controller) CreateCollectionItem(w http.ResponseWriter, r *http.Request) {
	const op = "CreateCollectionItem"
	log := c.log.With(slog.String("op", op))

	if err := r.ParseMultipartForm(10 << 20); err != nil {
		log.Error("failed to parse multipart form", "err", err)
		res.WriteError(w, "invalid form", http.StatusBadRequest)
		return
	}
	collectionID, err := strconv.Atoi(r.FormValue("collection_id"))
	if err != nil {
		res.WriteError(w, "invalid collection id", http.StatusBadRequest)
		return
	}
	file, _, err := r.FormFile("image")
	if err != nil {
		res.WriteError(w, "no image provided", http.StatusBadRequest)
		return
	}
	defer file.Close()
	payload := uc.CreateItemPayload{
		CollectionID: collectionID,
		Title:        r.FormValue("title"),
		Description:  r.FormValue("description"),
	}
	if err := c.validate.Struct(payload); err != nil {
		res.WriteError(w, "invalid payload", http.StatusBadRequest)
		return
	}
	item, err := c.usecase.Create(r.Context(), payload, file)
	if err != nil {
		if errors.Is(err, uc.ErrCollectionNotFound) {
			res.WriteError(w, "collection not found", http.StatusNotFound)
			return
		}
		res.WriteError(w, "failed to create collection item", http.StatusInternalServerError)
		return
	}
	res.WriteJSON(w, NewItem(item), http.StatusCreated)
}

// @Summary      Обновление украшения
// @Tags         items
// @Accept       multipart/form-data
// @Produce      json
// @Param        id            path      int     true  "ID украшения"
// @Param        title         formData  string  false "Новое название"
// @Param        description   formData  string  false "Новое описание"
// @Param        image         formData  file    false "Новое изображение"
// @Success      201  {object}  Item  "Item successfully updated"
// @Failure      400  {object}  res.ErrorResponse    "Invalid form data"
// @Failure      401  {object}  res.ErrorResponse  "Нет доступа"
// @Failure      404  {object}  res.ErrorResponse    "Item not found"
// @Failure      500  {object}  res.ErrorResponse    "Internal server error"
// @Router       /items/update/{id} [put]
func (c *controller) UpdateCollectionItem(w http.ResponseWriter, r *http.Request) {
	const op = "UpdateCollectionItem"
	log := c.log.With(slog.String("op", op))

	if err := r.ParseMultipartForm(10 << 20); err != nil {
		log.Error("failed to parse multipart form", "err", err)
		res.WriteError(w, "invalid form", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		res.WriteError(w, "invalid id", http.StatusBadRequest)
		return
	}
	title := r.FormValue("title")
	description := r.FormValue("description")
	payload := uc.UpdateItemPayload{ID: id}
	if title != "" {
		payload.Title = &title
	}
	if description != "" {
		payload.Description = &description
	}

	var image multipart.File
	if image, _, err = r.FormFile("image"); err == nil {
		defer image.Close()
	}

	if err := c.validate.Struct(payload); err != nil {
		res.WriteError(w, "invalid payload", http.StatusBadRequest)
		return
	}
	item, err := c.usecase.Update(r.Context(), payload, image)
	if err != nil {
		if errors.Is(err, uc.ErrItemNotFound) {
			res.WriteError(w, "collection item not found", http.StatusNotFound)
			return
		}
		res.WriteError(w, "failed to update collection item", http.StatusInternalServerError)
		return
	}
	res.WriteJSON(w, NewItem(item), http.StatusCreated)
}

// @Summary      Удаление украшения
// @Tags         items
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "ID украшения"
// @Success      200  {object}  Item  "Item successfully deleted"
// @Failure      400  {object}  res.ErrorResponse    "Invalid request parameters"
// @Failure      401  {object}  res.ErrorResponse  "Нет доступа"
// @Failure      404  {object}  res.ErrorResponse    "Item not found"
// @Failure      500  {object}  res.ErrorResponse    "Internal server error"
// @Router       /items/delete/{id} [delete]
func (c *controller) DeleteCollectionItem(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		res.WriteError(w, "invalid id", http.StatusBadRequest)
		return
	}
	item, err := c.usecase.Delete(r.Context(), id)
	if err != nil {
		if errors.Is(err, uc.ErrItemNotFound) {
			res.WriteError(w, "collection item not found", http.StatusNotFound)
			return
		}
		res.WriteError(w, "failed to delete collection item", http.StatusInternalServerError)
		return
	}
	res.WriteJSON(w, NewItem(item), http.StatusOK)
}

// @Summary      Получение украшений по ID коллекции
// @Description  Получение украшений по ID коллекции, сортировка по дате создания
// @Tags         items
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "ID коллекции"
// @Success      200  {array}   Item
// @Failure      400  {object}  res.ErrorResponse  "Invalid request parameters"
// @Failure      404  {object}  res.ErrorResponse  "Коллекция не найдена"
// @Failure      500  {object}  res.ErrorResponse  "Internal server error"
// @Router       /items/collection/{id} [get]
func (c *controller) ItemsByCollection(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		res.WriteError(w, "invalid id", http.StatusBadRequest)
		return
	}
	items, err := c.usecase.GetByCollectionId(r.Context(), id)
	if err != nil {
		if errors.Is(err, uc.ErrCollectionNotFound) {
			res.WriteError(w, "collection not found", http.StatusNotFound)
			return
		}
		res.WriteError(w, "failed to get collection items", http.StatusInternalServerError)
		return
	}
	result := make([]Item, len(items))
	for i, item := range items {
		result[i] = NewItem(&item)
	}
	res.WriteJSON(w, result, http.StatusOK)
}

// @Summary      Получение украшения по ID
// @Tags         items
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "ID украшения"
// @Success      200  {object}  Item
// @Failure      400  {object}  res.ErrorResponse  "Invalid request parameters"
// @Failure      404  {object}  res.ErrorResponse  "Украшние не найдено"
// @Failure      500  {object}  res.ErrorResponse  "Internal server error"
// @Router       /items/{id} [get]
func (c *controller) ById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		res.WriteError(w, "invalid id", http.StatusBadRequest)
		return
	}
	item, err := c.usecase.GetById(r.Context(), id)
	if err != nil {
		if errors.Is(err, uc.ErrItemNotFound) {
			res.WriteError(w, "collection item not found", http.StatusNotFound)
			return
		}
		res.WriteError(w, "failed to get collection item", http.StatusInternalServerError)
		return
	}
	res.WriteJSON(w, NewItem(item), http.StatusOK)
}
