package api

import (
	"context"
	"net/http"
	"time"

	"github.com/Jamshid90/flight/api/handlers/v1/flight"
	"github.com/Jamshid90/flight/internal/entity"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
)

type Option struct {
	Logger         *zap.Logger
	FlightUsecase  entity.FlightUsecase
	ContextTimeout time.Duration
}

//New ...
func New(option Option) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(option.ContextTimeout))
	r.Route("/v1", func(r chi.Router) {
		r.Use(apiVersionCtx("v1"))
		// flight handlers initialization
		r.Mount("/flight", flight.NewHandler(&flight.HandlerOption{Logger: option.Logger, FlightUsecase: option.FlightUsecase}))
	})
	return r
}

func apiVersionCtx(version string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			r = r.WithContext(context.WithValue(r.Context(), "api.version", version))
			next.ServeHTTP(w, r)
		})
	}
}
