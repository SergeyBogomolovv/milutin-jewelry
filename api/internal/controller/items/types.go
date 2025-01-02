package itemscontroller

import (
	"context"
	"mime/multipart"
	"time"

	storage "github.com/SergeyBogomolovv/milutin-jewelry/internal/storage/items"
	uc "github.com/SergeyBogomolovv/milutin-jewelry/internal/usecase/items"
)

type Usecase interface {
	Create(ctx context.Context, payload uc.CreateItemPayload, image multipart.File) (*storage.Item, error)
	Update(ctx context.Context, payload uc.UpdateItemPayload, image multipart.File) (*storage.Item, error)
	Delete(ctx context.Context, id int) (*storage.Item, error)
	GetByCollectionId(ctx context.Context, id int) ([]*storage.Item, error)
}

type Item struct {
	ID           int       `json:"id"`
	Title        string    `json:"title,omitempty"`
	Description  string    `json:"description,omitempty"`
	CollectionID int       `json:"collection_id"`
	ImageID      string    `json:"image_id"`
	CreatedAt    time.Time `json:"created_at"`
}

func NewItem(item *storage.Item) *Item {
	return &Item{
		ID:           item.ID,
		Title:        item.Title,
		Description:  item.Description,
		CollectionID: item.CollectionID,
		ImageID:      item.ImageID,
		CreatedAt:    item.CreatedAt,
	}
}
