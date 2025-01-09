package banner

import (
	"context"
	"mime/multipart"

	storage "github.com/SergeyBogomolovv/milutin-jewelry/internal/storage/banner"
)

type Storage interface {
	Save(ctx context.Context, collection *storage.Banner) error
	Delete(ctx context.Context, id int) error
	List(ctx context.Context) ([]storage.Banner, error)
	GetById(ctx context.Context, id int) (*storage.Banner, error)
	CollectionExists(ctx context.Context, id int) (bool, error)
}

type FilesService interface {
	UploadImage(ctx context.Context, image multipart.File, path string) (string, error)
	DeleteImage(ctx context.Context, key string) error
}

type CreateBannerPayload struct {
	CollectionID *int           `json:"collection_id,omitempty" example:"12"`
	Image        multipart.File `validate:"required" json:"image" example:"image.jpg"`
	MobileImage  multipart.File `validate:"required" json:"mobile_image" example:"image.jpg"`
}

const bannersKey = "banners"
