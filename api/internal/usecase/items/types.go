package itemsusecase

import (
	"context"
	"errors"
	"mime/multipart"
	"time"

	storage "github.com/SergeyBogomolovv/milutin-jewelry/internal/storage/items"
)

type Storage interface {
	Save(ctx context.Context, item *storage.Item) error
	Update(ctx context.Context, item *storage.Item) error
	Delete(ctx context.Context, id int) error
	GetById(ctx context.Context, id int) (*storage.Item, error)
	GetByCollectionId(ctx context.Context, id int) ([]*storage.Item, error)
	CollectionExists(ctx context.Context, id int) (bool, error)
}

type FilesService interface {
	UploadImage(ctx context.Context, image multipart.File, dest string) (string, error)
	DeleteImage(ctx context.Context, key string) error
}

type Item struct {
	ID           int       `json:"id"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	CollectionID int       `json:"collection_id"`
	ImageID      string    `json:"image_id"`
	CreatedAt    time.Time `json:"created_at"`
}

func newItem(item *storage.Item) *Item {
	return &Item{
		ID:           item.ID,
		Title:        item.Title,
		Description:  item.Description,
		CollectionID: item.CollectionID,
		ImageID:      item.ImageID,
		CreatedAt:    item.CreatedAt,
	}
}

var (
	ErrCollectionNotFound = errors.New("collection not exist")
	ErrItemNotFound       = errors.New("item not found")
)

const itemsKey = "items"
