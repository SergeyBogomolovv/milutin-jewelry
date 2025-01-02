package code

import "errors"

var (
	ErrInvalidCode  = errors.New("invalid otp code")
	ErrCodeNotFound = errors.New("code not found")
)
