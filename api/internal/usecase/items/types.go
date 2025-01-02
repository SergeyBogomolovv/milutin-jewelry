package itemsusecase

import (
	"context"
	"errors"
	"mime/multipart"

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

type CreateItemPayload struct {
	CollectionID int
	Title        string
	Description  string
}

type UpdateItemPayload struct {
	ID          int
	Title       *string
	Description *string
}

var (
	ErrCollectionNotFound = errors.New("collection not exist")
	ErrItemNotFound       = errors.New("item not found")
)

const itemsKey = "items"
