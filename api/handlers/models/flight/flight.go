package flight

import (
	"time"

	"github.com/Jamshid90/flight/internal/entity"
)

type FlightJson struct {
	ID            int       `json:"id,omitempty"`
	Number        int       `json:"number,omitempty"`
	CityDeparture string    `json:"city_departure,omitempty"`
	CityArrival   string    `json:"city_arrival,omitempty"`
	TimeDeparture time.Time `json:"time_departure,omitempty"`
	TimeArrival   time.Time `json:"time_arrival,omitempty"`
}

func ToJson(flight *entity.Flight) *FlightJson {
	return &FlightJson{
		ID:            flight.ID,
		Number:        flight.Number,
		CityDeparture: flight.CityDeparture,
		CityArrival:   flight.CityArrival,
		TimeDeparture: flight.TimeDeparture,
		TimeArrival:   flight.TimeArrival,
	}
}

func ToJsons(flights []*entity.Flight) []*FlightJson {
	var josnFlights []*FlightJson
	for _, flight := range flights {
		josnFlights = append(josnFlights, ToJson(flight))
	}
	return josnFlights
}
