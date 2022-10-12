package model

import (
	"errors"
)

var (
	ErrNotFound         = errors.New("not found")
	ErrSomethingStrange = errors.New("crushed data")
	ErrEmptyData        = errors.New("empty data")
)
