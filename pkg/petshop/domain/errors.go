package domain

import "errors"

var (
	ErrBadParamInput = errors.New("Given param is not valid")
	ErrNotFound      = errors.New("Not Found")
)
