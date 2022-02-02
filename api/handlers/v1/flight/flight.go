package flight

import (
	"encoding/json"
	"io"
	"net/http"

	errapi "github.com/Jamshid90/flight/api/errors"
	flightModel "github.com/Jamshid90/flight/api/handlers/models/flight"
	"github.com/Jamshid90/flight/internal/entity"
	"github.com/Jamshid90/flight/internal/validation"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"go.uber.org/zap"
)

type HandlerOption struct {
	Logger        *zap.Logger
	FlightUsecase entity.FlightUsecase
}

type flightHandler struct {
	logger        *zap.Logger
	flightUsecase entity.FlightUsecase
}

// new handler ...
func NewHandler(option *HandlerOption) chi.Router {
	handler := flightHandler{
		logger:        option.Logger,
		flightUsecase: option.FlightUsecase,
	}

	r := chi.NewRouter()
	r.Group(func(r chi.Router) {
		r.Post("/", handler.create())
		r.Get("/", handler.list())
	})
	return r
}

func (handler *flightHandler) create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			createRequest flightModel.CreateRequest
		)

		requestBody, err := io.ReadAll(r.Body)
		if err != nil {
			render.Render(w, r, errapi.ErrInvalidArgument)
			return
		}
		defer r.Body.Close()

		if json.Unmarshal(requestBody, &createRequest); err != nil {
			render.Render(w, r, errapi.ErrInvalidArgument)
			return
		}

		if err := validation.Validator(&createRequest); err != nil {
			render.Render(w, r, errapi.ErrValidatorRender(err))
			return
		}

		flights, err := flightModel.FlightsToEntity(createRequest.Flights)

		if err != nil {
			render.Render(w, r, errapi.ErrInvalidArgument)
			return
		}

		if err := handler.flightUsecase.CreateBulk(r.Context(), flights); err != nil {
			render.Render(w, r, errapi.ErrInvalidRequestRender(err))
			return
		}

		render.JSON(w, r, 200)
	}
}

func (handler *flightHandler) list() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		list, err := handler.flightUsecase.List(r.Context(), r.URL.Query())
		if err != nil {
			render.Render(w, r, errapi.ErrInvalidArgument)
			return
		}
		render.JSON(w, r, flightModel.ToJsons(list))
	}
}
