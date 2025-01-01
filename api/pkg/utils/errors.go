package utils

import "fmt"

func WrapErr(msg string, err error) error {
	return fmt.Errorf("%s: %w", msg, err)
}

func WrapIfErr(msg string, err error) error {
	if err == nil {
		return nil
	}
	return WrapErr(msg, err)
}
