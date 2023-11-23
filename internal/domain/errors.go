package domain

import "errors"

var (
	ErrIDMismatch = errors.New("ID mismatch between URL and JSON data")
)
