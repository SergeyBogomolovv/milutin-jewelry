package collections

import (
	"context"
	"mime/multipart"

	storage "github.com/SergeyBogomolovv/milutin-jewelry/internal/storage/collection"
	uc "github.com/SergeyBogomolovv/milutin-jewelry/internal/usecase/collection"
)

type Usecase interface {
	Create(ctx context.Context, payload uc.CreateCollectionPayload, image multipart.File) (*storage.Collection, error)
	Update(ctx context.Context, payload uc.UpdateCollectionPayload, image multipart.File) (*storage.Collection, error)
	Delete(ctx context.Context, id int) (*storage.Collection, error)
	GetAll(ctx context.Context) ([]storage.Collection, error)
	GetByID(ctx context.Context, id int) (*storage.Collection, error)
}

type Collection struct {
	ID          int    `json:"id" example:"1"`
	Title       string `json:"title" example:"Кольца"`
	Description string `json:"description" example:"Описание коллекции"`
	ImageID     string `json:"image_id" example:"12345"`
	CreatedAt   string `json:"created_at" example:"2025-01-01T12:00:00Z"`
}

func NewCollection(collection *storage.Collection) Collection {
	return Collection{
		ID:          collection.ID,
		Title:       collection.Title,
		Description: collection.Description,
		ImageID:     collection.ImageID,
		CreatedAt:   collection.CreatedAt,
	}
}
