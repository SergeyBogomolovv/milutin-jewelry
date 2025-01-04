package collectionscontroller

import (
	"errors"
	"fmt"
	"log/slog"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/SergeyBogomolovv/milutin-jewelry/internal/middleware"
	uc "github.com/SergeyBogomolovv/milutin-jewelry/internal/usecase/collection"
	"github.com/SergeyBogomolovv/milutin-jewelry/pkg/lib/res"
	"github.com/go-playground/validator/v10"
)

type controller struct {
	validate *validator.Validate
	usecase  Usecase
	log      *slog.Logger
}

func Register(log *slog.Logger, router *http.ServeMux, usecase Usecase, auth middleware.Middleware) {
	const dest = "collectionsController"
	c := &controller{
		log:      log.With(slog.String("dest", dest)),
		usecase:  usecase,
		validate: validator.New(validator.WithRequiredStructEnabled()),
	}
	r := http.NewServeMux()
	r.HandleFunc("GET /all", c.GetAllCollections)
	r.HandleFunc("GET /{id}", c.GetOneCollection)

	pr := http.NewServeMux()
	pr.HandleFunc("POST /create", c.CreateCollection)
	pr.HandleFunc("PUT /update/{id}", c.UpdateCollection)
	pr.HandleFunc("DELETE /delete/{id}", c.DeleteCollection)
	r.Handle("/", auth(pr))

	router.Handle("/collections/", http.StripPrefix("/collections", r))
}

// @Summary      Создание коллекции
// @Description  Создание новой коллекции
// @Tags         collections
// @Accept       multipart/form-data
// @Produce      json
// @Param        title        formData  string  true   "Название коллекции"
// @Param        description  formData  string  false  "Описание коллекции"
// @Param        image        formData  file    true   "Изображение коллекции"
// @Success      201  {object}  Collection  "Коллекция успешно создана"
// @Failure      400  {object}  res.ErrorResponse  "Неверный формат данных"
// @Failure      401  {object}  res.ErrorResponse  "Нет доступа"
// @Failure      500  {object}  res.ErrorResponse  "Ошибка сервера"
// @Router       /collections/create [post]
func (c *controller) CreateCollection(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		res.WriteError(w, "invalid payload", http.StatusBadRequest)
		return
	}
	image, _, err := r.FormFile("image")
	if err != nil {
		res.WriteError(w, "no image provided", http.StatusBadRequest)
		return
	}
	defer image.Close()

	payload := uc.CreateCollectionPayload{Title: r.FormValue("title"), Description: r.FormValue("description")}
	if err := c.validate.Struct(payload); err != nil {
		c.log.Error("failed to validate payload", "err", err)
		res.WriteError(w, fmt.Sprintf("invalid payload: %s", err), http.StatusBadRequest)
		return
	}

	collection, err := c.usecase.Create(r.Context(), payload, image)
	if err != nil {
		res.WriteError(w, "failed to create collection", http.StatusInternalServerError)
		return
	}
	res.WriteJSON(w, NewCollection(collection), http.StatusCreated)
}

// @Summary      Обновление коллекции
// @Description  Обновление данных существующей коллекции
// @Tags         collections
// @Accept       multipart/form-data
// @Produce      json
// @Param        id           path      int     true   "ID коллекции"
// @Param        title        formData  string  false  "Новое название коллекции"
// @Param        description  formData  string  false  "Новое описание коллекции"
// @Param        image        formData  file    false  "Новое изображение коллекции"
// @Success      200  {object}  Collection  "Коллекция успешно обновлена"
// @Failure      400  {object}  res.ErrorResponse  "Неверный формат данных"
// @Failure      401  {object}  res.ErrorResponse  "Нет доступа"
// @Failure      404  {object}  res.ErrorResponse  "Коллекция не найдена"
// @Failure      500  {object}  res.ErrorResponse  "Ошибка сервера"
// @Router       /collections/update/{id} [put]
func (c *controller) UpdateCollection(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		res.WriteError(w, "invalid payload", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		res.WriteError(w, "invalid id", http.StatusBadRequest)
		return
	}
	payload := uc.UpdateCollectionPayload{ID: id}
	title := r.FormValue("title")
	description := r.FormValue("description")
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
	collection, err := c.usecase.Update(r.Context(), payload, image)
	if err != nil {
		if errors.Is(err, uc.ErrCollectionNotFound) {
			res.WriteError(w, "collection not found", http.StatusNotFound)
			return
		}
		res.WriteError(w, "failed to update collection", http.StatusInternalServerError)
		return
	}
	res.WriteJSON(w, NewCollection(collection), http.StatusCreated)
}

// @Summary      Удаление коллекции
// @Description  Удаление существующей коллекции
// @Tags         collections
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "ID коллекции"
// @Success      200  {object}  Collection  "Коллекция успешно удалена"
// @Failure      400  {object}  res.ErrorResponse  "Неверный формат данных"
// @Failure      401  {object}  res.ErrorResponse  "Нет доступа"
// @Failure      404  {object}  res.ErrorResponse  "Коллекция не найдена"
// @Failure      500  {object}  res.ErrorResponse  "Ошибка сервера"
// @Router       /collections/delete/{id} [delete]
func (c *controller) DeleteCollection(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		res.WriteError(w, "invalid id", http.StatusBadRequest)
		return
	}
	collection, err := c.usecase.Delete(r.Context(), id)
	if err != nil {
		if errors.Is(err, uc.ErrCollectionNotFound) {
			res.WriteError(w, "collection not found", http.StatusNotFound)
			return
		}
		res.WriteError(w, "failed to delete collection", http.StatusInternalServerError)
		return
	}
	res.WriteJSON(w, NewCollection(collection), http.StatusOK)
}

// @Summary      Получение всех коллекций
// @Description  Получение списка всех коллекций
// @Tags         collections
// @Accept       json
// @Produce      json
// @Success      200  {array}  Collection  "Список всех коллекций"
// @Failure      500  {object}  res.ErrorResponse  "Ошибка сервера"
// @Router       /collections/all [get]
func (c *controller) GetAllCollections(w http.ResponseWriter, r *http.Request) {
	collections, err := c.usecase.GetAll(r.Context())
	if err != nil {
		res.WriteError(w, "failed to get collections", http.StatusInternalServerError)
		return
	}
	result := make([]Collection, len(collections))
	for i, collection := range collections {
		result[i] = NewCollection(&collection)
	}
	res.WriteJSON(w, result, http.StatusOK)
}

// @Summary      Получение одной коллекции
// @Description  Получение данных одной коллекции по её ID
// @Tags         collections
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "ID коллекции"
// @Success      200  {object}  Collection  "Данные коллекции"
// @Failure      400  {object}  res.ErrorResponse  "Неверный формат данных"
// @Failure      404  {object}  res.ErrorResponse  "Коллекция не найдена"
// @Failure      500  {object}  res.ErrorResponse  "Ошибка сервера"
// @Router       /collections/{id} [get]
func (c *controller) GetOneCollection(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		res.WriteError(w, "invalid id", http.StatusBadRequest)
		return
	}
	collection, err := c.usecase.GetByID(r.Context(), id)
	if err != nil {
		if errors.Is(err, uc.ErrCollectionNotFound) {
			res.WriteError(w, "collection not found", http.StatusNotFound)
			return
		}
		res.WriteError(w, "failed to get collection", http.StatusInternalServerError)
		return
	}
	res.WriteJSON(w, NewCollection(collection), http.StatusOK)
}
