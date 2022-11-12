package entity

import (
	"time"
)

type Flight struct {
	Id   id.Id `json:"id"`
	View FlightView
}

var _ Entity = (*Flight)(nil)

type FlightView struct {
	AirlCode string `json:"airlCode"`

	FltNum  string `json:"fltNum"`
	FltDate string `json:"fltDate"`

	OrigIATA string `json:"origIata"`
	DestIATA string `json:"destIata"`

	DepartureTime time.Time `json:"departureTime"`
	ArriveTime    time.Time `json:"arriveTime"`

	TotalCash float64 `json:"totalCash"`

	CorrectlyParsed bool `json:"correctlyParsed"`
}

func FromFlightView(id id.Id, view FlightView) *Flight {
	return &Flight{
		Id:   id,
		View: view,
	}
}

func ToFlightView(f Flight) FlightView {
	return f.View
}
