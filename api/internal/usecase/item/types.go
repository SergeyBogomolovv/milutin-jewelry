package item

import (
	"context"
	"errors"
	"mime/multipart"

	storage "github.com/SergeyBogomolovv/milutin-jewelry/internal/storage/item"
)

type Storage interface {
	Save(ctx context.Context, item *storage.Item) error
	Update(ctx context.Context, item *storage.Item) error
	Delete(ctx context.Context, id int) error
	GetById(ctx context.Context, id int) (*storage.Item, error)
	GetByCollectionId(ctx context.Context, id int) ([]storage.Item, error)
	CollectionExists(ctx context.Context, id int) (bool, error)
}

type FilesService interface {
	UploadImage(ctx context.Context, image multipart.File, dest string) (string, error)
	DeleteImage(ctx context.Context, key string) error
}

type CreateItemPayload struct {
	CollectionID int    `validate:"required" json:"collection_id" example:"1"`
	Title        string `json:"title" example:"Кулон"`
	Description  string `json:"description" example:"Описание кулона"`
}

type UpdateItemPayload struct {
	ID          int     `validate:"required" json:"id" example:"1"`
	Title       *string `json:"title,omitempty" example:"Серьги"`
	Description *string `json:"description,omitempty" example:"Новое описание украшения"`
}

var (
	ErrCollectionNotFound = errors.New("collection not exist")
	ErrItemNotFound       = errors.New("item not found")
)

const itemsKey = "items"
