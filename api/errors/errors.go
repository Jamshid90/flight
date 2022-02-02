package errors

import (
	"errors"
	"net/http"

	errapi "github.com/Jamshid90/flight/internal/errors"
	"github.com/go-chi/render"
)

var (
	ErrInvalidArgument = ErrInvalidRequestRender(errors.New("invalid argument"))
	ErrNotFound        = &ErrResponse{HTTPStatusCode: 404, ErrorText: "resource not found."}
	errValidation      *errapi.ErrValidation
)

type ErrResponse struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	AppCode   int64             `json:"code,omitempty"`  // application-specific error code
	ErrorText string            `json:"error,omitempty"` // application-level error message, for debugging
	Errors    map[string]string `json:"errors,omitempty"`
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func ErrInvalidRequestRender(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 400,
		ErrorText:      err.Error(),
	}
}

func ErrInvalidArgumentRender(err error, errors map[string]string) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 422,
		ErrorText:      err.Error(),
		Errors:         errors,
	}
}

func ErrValidatorRender(err error) render.Renderer {
	if errors.As(err, &errValidation) {
		return &ErrResponse{
			Err:            err,
			HTTPStatusCode: 422,
			ErrorText:      err.Error(),
			Errors:         errValidation.Errors,
		}
	}
	return ErrInvalidArgument
}
