package errs

import "errors"

var (
	ErrInvalidLoginCode = errors.New("invalid login code")
)
