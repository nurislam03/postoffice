package errors

import "errors"

// List of defined errors
var (
	ErrServerClosed     = errors.New("server closed")
	ErrResourceNotFound = errors.New("resource not found")
	ErrBadRequestBody   = errors.New("bad request body")
	ErrUnauthorized     = errors.New("unauthorized")
	ErrForbidden        = errors.New("forbidden")
	ErrInvalidToken     = errors.New("invalid token")
	ErrUnknown          = errors.New("unknown error")
	ErrMethodNotAllowed = errors.New("method not allowd")
	ErrInvalidID        = errors.New("invalid id")
	ErrDuplicateKey     = errors.New("duplicate key")
)
