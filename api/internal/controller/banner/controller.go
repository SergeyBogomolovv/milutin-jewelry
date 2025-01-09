package banner

import (
	"errors"
	"net/http"
	"strconv"

	uc "github.com/SergeyBogomolovv/milutin-jewelry/internal/usecase/banner"

	"github.com/SergeyBogomolovv/milutin-jewelry/internal/middleware"
	"github.com/SergeyBogomolovv/milutin-jewelry/pkg/lib/res"
)

type controller struct {
	usecase Usecase
}

func Register(router *http.ServeMux, usecase Usecase, auth middleware.Middleware) {
	const dest = "bannerController"
	controller := &controller{
		usecase: usecase,
	}
	r := http.NewServeMux()
	r.HandleFunc("GET /all", controller.GetAll)

	pr := http.NewServeMux()
	pr.HandleFunc("POST /create", controller.Create)
	pr.HandleFunc("DELETE /delete/{id}", controller.Delete)

	r.Handle("/", auth(pr))

	router.Handle("/banners/", http.StripPrefix("/banners", r))
}

// @Summary      Получение всех баннеров
// @Tags         banners
// @Accept       json
// @Produce      json
// @Success      200  {array}  Banner  "Список всех баннеров"
// @Failure      500  {object}  res.ErrorResponse  "Ошибка сервера"
// @Router       /banners/all [get]
func (c *controller) GetAll(w http.ResponseWriter, r *http.Request) {
	stored, err := c.usecase.List(r.Context())
	if err != nil {
		res.WriteError(w, "failed to get banners", http.StatusInternalServerError)
		return
	}
	banners := make([]Banner, len(stored))
	for i, banner := range stored {
		banners[i] = NewBanner(banner)
	}
	res.WriteJSON(w, banners, http.StatusOK)
}

// @Summary      Создание баннера
// @Tags         banners
// @Accept       multipart/form-data
// @Produce      json
// @Param        collection_id  formData  string  false  "ID коллекции"
// @Param        image          formData  file    true   "Изображение"
// @Param        mobile_image   formData  file    true   "Изображение для мобильных устройств"
// @Success      201  {object}  Banner  "Баннер создан"
// @Failure      400  {object}  res.ErrorResponse  "Неверный формат данных"
// @Failure      401  {object}  res.ErrorResponse  "Нет доступа"
// @Failure      404  {object}  res.ErrorResponse  "Коллекция не найдена"
// @Failure      500  {object}  res.ErrorResponse  "Ошибка сервера"
// @Router       /banners/create [post]
func (c *controller) Create(w http.ResponseWriter, r *http.Request) {
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

	mobileImage, _, err := r.FormFile("mobile_image")
	if err != nil {
		res.WriteError(w, "no mobile image provided", http.StatusBadRequest)
		return
	}
	defer mobileImage.Close()

	payload := uc.CreateBannerPayload{
		Image:       image,
		MobileImage: mobileImage,
	}

	if r.Form.Has("collection_id") {
		collectionID, err := strconv.Atoi(r.FormValue("collection_id"))
		if err != nil {
			res.WriteError(w, "invalid collection_id", http.StatusBadRequest)
			return
		}
		payload.CollectionID = &collectionID
	}

	banner, err := c.usecase.Create(r.Context(), payload)
	if err != nil {
		if errors.Is(err, uc.ErrCollectionNotFound) {
			res.WriteError(w, "collection not found", http.StatusNotFound)
			return
		}
		res.WriteError(w, "failed to create banner", http.StatusInternalServerError)
		return
	}
	res.WriteJSON(w, NewBanner(*banner), http.StatusCreated)
}

// @Summary      Удаление баннера
// @Tags         banners
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "ID баннера"
// @Success      200  {object}  Banner  "Баннер успешно удален"
// @Failure      400  {object}  res.ErrorResponse  "Неверный формат данных"
// @Failure      401  {object}  res.ErrorResponse  "Нет доступа"
// @Failure      404  {object}  res.ErrorResponse  "Баннер не найден"
// @Failure      500  {object}  res.ErrorResponse  "Ошибка сервера"
// @Router       /banners/delete/{id} [delete]
func (c *controller) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		res.WriteError(w, "invalid id", http.StatusBadRequest)
		return
	}

	banner, err := c.usecase.Delete(r.Context(), id)
	if err != nil {
		if errors.Is(err, uc.ErrBannerNotFound) {
			res.WriteError(w, "banner not found", http.StatusNotFound)
			return
		}
		res.WriteError(w, "failed to delete banner", http.StatusInternalServerError)
		return
	}
	res.WriteJSON(w, NewBanner(*banner), http.StatusOK)
}
