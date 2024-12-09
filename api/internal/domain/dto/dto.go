package dto

type LoginDTO struct {
	Code string `json:"code" validate:"len=6"`
}
