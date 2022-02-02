package flight

import (
	"time"

	"github.com/Jamshid90/flight/internal/entity"
)

type FlightValidate struct {
	Number        int    `json:"number" validate:"required"`
	CityDeparture string `json:"city_departure" validate:"required"`
	CityArrival   string `json:"city_arrival" validate:"required"`
	TimeDeparture string `json:"time_departure" validate:"required,datetime=2006-01-02T15:04:05Z07:00"`
	TimeArrival   string `json:"time_arrival" validate:"required,datetime=2006-01-02T15:04:05Z07:00"`
}

type CreateRequest struct {
	Flights []*FlightValidate `json:"flights" validate:"required,dive,required"`
}

func FlightToEntity(flight *FlightValidate) (*entity.Flight, error) {
	timeDeparture, err := time.Parse(time.RFC3339, flight.TimeDeparture)
	if err != nil {
		return nil, err
	}

	timeArrival, err := time.Parse(time.RFC3339, flight.TimeArrival)
	if err != nil {
		return nil, err
	}

	return &entity.Flight{
		Number:        flight.Number,
		CityDeparture: flight.CityDeparture,
		CityArrival:   flight.CityArrival,
		TimeDeparture: timeDeparture,
		TimeArrival:   timeArrival,
	}, nil
}

func FlightsToEntity(flights []*FlightValidate) ([]*entity.Flight, error) {
	var entityFlights []*entity.Flight
	for _, flight := range flights {
		f, err := FlightToEntity(flight)
		if err != nil {
			return entityFlights, err
		}
		entityFlights = append(entityFlights, f)
	}
	return entityFlights, nil
}
