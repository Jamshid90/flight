package flight

import (
	"context"
	"time"

	"github.com/Jamshid90/flight/internal/entity"
)

type flightUseCase struct {
	contextTimeout time.Duration
	flightRepo     entity.FlightRepository
}

func NewUsecase(timeout time.Duration, flightRepo entity.FlightRepository) entity.FlightUsecase {
	return &flightUseCase{
		contextTimeout: timeout,
		flightRepo:     flightRepo,
	}
}

func (u *flightUseCase) beforeCreate(flight *entity.Flight) error {
	flight.CreatedAt = time.Now().UTC()
	flight.UpdatedAt = time.Now().UTC()
	return nil
}

func (u *flightUseCase) Create(ctx context.Context, flight *entity.Flight) error {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	u.beforeCreate(flight)
	return u.flightRepo.Create(ctx, flight)
}

func (u *flightUseCase) CreateBulk(ctx context.Context, flights []*entity.Flight) error {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()
	return u.flightRepo.CreateBulk(ctx, flights)
}

func (u *flightUseCase) Read(ctx context.Context, id int64) (*entity.Flight, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()
	return u.flightRepo.Read(ctx, id)
}

func (u *flightUseCase) beforeUpdate(flight *entity.Flight) error {
	flight.UpdatedAt = time.Now().UTC()
	return nil
}

func (u *flightUseCase) Update(ctx context.Context, flight *entity.Flight) error {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	u.beforeUpdate(flight)
	return u.flightRepo.Update(ctx, flight)
}

func (u *flightUseCase) Delete(ctx context.Context, id int64) error {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()
	return u.flightRepo.Delete(ctx, id)
}

func (u *flightUseCase) List(ctx context.Context, parameters map[string][]string) ([]*entity.Flight, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()
	return u.flightRepo.List(ctx, parameters)
}
