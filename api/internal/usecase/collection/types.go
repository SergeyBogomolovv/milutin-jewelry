package collection

import (
	"context"
	"errors"
	"mime/multipart"

	storage "github.com/SergeyBogomolovv/milutin-jewelry/internal/storage/collection"
)

type Storage interface {
	Save(ctx context.Context, collection *storage.Collection) error
	Update(ctx context.Context, collection *storage.Collection) error
	Delete(ctx context.Context, id int) error
	GetByID(ctx context.Context, id int) (*storage.Collection, error)
	GetAll(ctx context.Context) ([]storage.Collection, error)
}

type FilesService interface {
	UploadImage(ctx context.Context, image multipart.File, path string) (string, error)
	DeleteImage(ctx context.Context, key string) error
}

type CreateCollectionPayload struct {
	Title       string `validate:"required" json:"title" example:"Кольца"`
	Description string `json:"description" example:"Описание коллекции"`
}

type UpdateCollectionPayload struct {
	ID          int     `validate:"required" json:"id" example:"1"`
	Title       *string `json:"title,omitempty" example:"Кулоны"`
	Description *string `json:"description,omitempty" example:"Новое описание коллекции"`
}

const collectionsKey = "collections"

var (
	ErrCollectionNotFound = errors.New("collection not found")
)
