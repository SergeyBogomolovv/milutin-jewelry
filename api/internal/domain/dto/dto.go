package dto

import "mime/multipart"

type LoginDTO struct {
	Code string `json:"code" validate:"len=6"`
}

type CreateCollectionDTO struct {
	Title       string                `json:"title" validate:"required"`
	Description string                `json:"description"`
	Image       *multipart.FileHeader `json:"image"`
}

type UpdateCollectionDTO struct {
	Title       string                `json:"title"`
	Description string                `json:"description"`
	Image       *multipart.FileHeader `json:"image"`
}

// TODO: update
type CollectionDTO struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       string `json:"image"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
