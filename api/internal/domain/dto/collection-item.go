package dto

import "mime/multipart"

type CreateCollectionItemRequest struct {
	Title        string `validate:"required"`
	Description  string
	Image        multipart.File `validate:"required"`
	CollectionID int            `validate:"required"`
}

type CreateCollectionItemInput struct {
	Title        string
	Description  string
	ImageID      string
	CollectionID int
}

type UpdateCollectionItemRequest struct {
	Title       string
	Description string
	Image       multipart.File
	ID          int
}

type UpdateCollectionItemInput struct {
	Title       string
	Description string
	ImageID     string
	ID          int
}

type CollectionItemResponse struct {
	ID           int    `json:"id"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	ImageID      string `json:"image_id"`
	CollectionID int    `json:"collection_id"`
}
