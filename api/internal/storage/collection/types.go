package collection

import "errors"

type Collection struct {
	ID          int    `db:"collection_id"`
	Title       string `db:"title"`
	Description string `db:"description"`
	ImageID     string `db:"image_id"`
	CreatedAt   string `db:"created_at"`
}

var (
	ErrCollectionNotFound = errors.New("collection not found")
)
