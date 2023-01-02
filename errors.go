package cache

import "errors"

var (
	ErrNotFound = errors.New("value by provided key not found")
)
