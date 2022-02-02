package errors

import (
	"errors"
	"net/http"
	"strings"
)

var (
	BadRequest        = NewErrNotFound(GetHTTPStatusText(http.StatusBadRequest))
	ErrInternalServer = errors.New(GetHTTPStatusText(http.StatusInternalServerError))
	ErrorNotFound     = NewErrNotFound("object")
)

// GetHTTPStatusText ...
func GetHTTPStatusText(statusCode int) string {
	return strings.ToLower(http.StatusText(statusCode))
}

// error not found
func NewErrNotFound(text string) *ErrNotFound {
	return &ErrNotFound{text}
}

type ErrNotFound struct {
	name string
}

func (e *ErrNotFound) Error() string {
	return e.name + " not found"
}

// error bad request
func NewErrBadRequest(err error, message string) *ErrBadRequest {
	return &ErrBadRequest{err, message}
}

type ErrBadRequest struct {
	Err     error
	Message string
}

func (e ErrBadRequest) Error() string {
	return e.Message
}

// error validation
func NewErrValidation() *ErrValidation {
	return &ErrValidation{Errors: make(map[string]string)}
}

type ErrValidation struct {
	Err    error
	Errors map[string]string
}

func (e ErrValidation) Error() string {
	return e.Err.Error()
}
