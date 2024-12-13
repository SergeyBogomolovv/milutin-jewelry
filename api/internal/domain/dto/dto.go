package dto

import "mime/multipart"

type LoginDTO struct {
	Code string `json:"code" validate:"len=6"`
}

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
	ImageHeader *multipart.FileHeader
}

type UpdateCollectionInput struct {
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
