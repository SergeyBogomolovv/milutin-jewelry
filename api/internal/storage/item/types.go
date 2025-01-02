package item

import (
	"errors"
	"time"
)

type Item struct {
	ID           int       `db:"item_id"`
	CollectionID int       `db:"collection_id"`
	Title        string    `db:"title"`
	Description  string    `db:"description"`
	ImageID      string    `db:"image_id"`
	CreatedAt    time.Time `db:"created_at"`
}

var (
	ErrItemNotFound = errors.New("item not found")
)
