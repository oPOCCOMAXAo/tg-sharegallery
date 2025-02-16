package models

import "errors"

var (
	ErrFailed   = errors.New("failed")
	ErrNotFound = errors.New("not found")
)
