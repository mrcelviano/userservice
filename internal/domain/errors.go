package domain

import "errors"

var (
	ErrInvalidEmail        = errors.New("email is invalid")
	ErrInternalServerError = errors.New("internal server error")
	ErrNotFound            = errors.New("your requested item is not found")
	ErrBadParamInput       = errors.New("email or name already exists")
)
