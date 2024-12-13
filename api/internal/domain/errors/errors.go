package errs

import "errors"

var (
	ErrInvalidLoginCode   = errors.New("invalid login code")
	ErrCollectionNotFound = errors.New("collection not found")
)
