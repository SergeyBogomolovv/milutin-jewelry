package dto

import "mime/multipart"

type CreateCollectionRequest struct {
	Title       string `validate:"required"`
	Description string
	Image       multipart.File `validate:"required"`
}

type CreateCollectionInput struct {
	Title       string
	Description string
	ImageID     string
}

type UpdateCollectionRequest struct {
	Title       string
	Description string
	Image       multipart.File
	ID          int
}

type UpdateCollectionInput struct {
	ID          int
	Title       string
	Description string
	ImageID     string
}

type CollectionResponse struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ImageID     string `json:"image_id"`
}
