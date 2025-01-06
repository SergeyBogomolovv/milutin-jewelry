package items

import (
	"context"
	"mime/multipart"
	"time"

	storage "github.com/SergeyBogomolovv/milutin-jewelry/internal/storage/item"
	uc "github.com/SergeyBogomolovv/milutin-jewelry/internal/usecase/item"
)

type Usecase interface {
	Create(ctx context.Context, payload uc.CreateItemPayload, image multipart.File) (*storage.Item, error)
	Update(ctx context.Context, payload uc.UpdateItemPayload, image multipart.File) (*storage.Item, error)
	Delete(ctx context.Context, id int) (*storage.Item, error)
	GetByCollectionId(ctx context.Context, id int) ([]storage.Item, error)
	GetById(ctx context.Context, id int) (*storage.Item, error)
}

type Item struct {
	ID           int       `json:"id" example:"1"`
	Title        string    `json:"title,omitempty" example:"Кольцо"`
	Description  string    `json:"description,omitempty" example:"Описание украшения"`
	CollectionID int       `json:"collection_id" example:"1"`
	ImageID      string    `json:"image_id" example:"67890"`
	CreatedAt    time.Time `json:"created_at" example:"2025-01-01T12:00:00Z"`
}

func NewItem(item *storage.Item) Item {
	return Item{
		ID:           item.ID,
		Title:        item.Title,
		Description:  item.Description,
		CollectionID: item.CollectionID,
		ImageID:      item.ImageID,
		CreatedAt:    item.CreatedAt,
	}
}
