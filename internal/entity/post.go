package entity

import (
	"context"
	"time"
)

type Flight struct {
	ID            int
	Number        int
	CityDeparture string
	CityArrival   string
	TimeDeparture time.Time
	TimeArrival   time.Time
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type FlightRepository interface {
	Create(ctx context.Context, Flight *Flight) error
	CreateBulk(ctx context.Context, flights []*Flight) error
	Read(ctx context.Context, id int64) (*Flight, error)
	Update(ctx context.Context, Flight *Flight) error
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context, parameters map[string][]string) ([]*Flight, error)
}

type FlightUsecase interface {
	Create(ctx context.Context, Flight *Flight) error
	CreateBulk(ctx context.Context, flights []*Flight) error
	Read(ctx context.Context, id int64) (*Flight, error)
	Update(ctx context.Context, Flight *Flight) error
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context, parameters map[string][]string) ([]*Flight, error)
}
