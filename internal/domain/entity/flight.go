package entity

import (
	"api-app/pkg/object/oid"
)

type Flight struct {
	Entity `json:"-" db:"-"`
	Id     oid.Id `json:"id" db:"id"`

	// View
	AirlCode        string  `json:"airlCode" db:"airl_code"`
	FltNum          string  `json:"fltNum" db:"flt_num"`
	FltDate         string  `json:"fltDate" db:"flt_date"`
	OrigIATA        string  `json:"origIata" db:"orig_iata"`
	DestIATA        string  `json:"destIata" db:"dest_iata"`
	DepartureTime   string  `json:"departureTime" db:"departure_time"`
	ArrivalTime     string  `json:"arrivalTime" db:"arrival_time"`
	TotalCash       float64 `json:"totalCash" db:"total_cash"`
	CorrectlyParsed bool    `json:"correctlyParsed" db:"correctly_parsed"`
}

func (f *Flight) ToView() *FlightView {
	return &FlightView{
		AirlCode:        f.AirlCode,
		FltNum:          f.FltNum,
		FltDate:         f.FltDate,
		OrigIATA:        f.OrigIATA,
		DestIATA:        f.DestIATA,
		DepartureTime:   f.DepartureTime,
		ArrivalTime:     f.ArrivalTime,
		TotalCash:       f.TotalCash,
		CorrectlyParsed: f.CorrectlyParsed,
	}
}

type FlightView struct {
	View            `json:"-" db:"-"`
	AirlCode        string  `json:"airlCode" db:"airl_code"`
	FltNum          string  `json:"fltNum" db:"flt_num"`
	FltDate         string  `json:"fltDate" db:"flt_date"`
	OrigIATA        string  `json:"origIata" db:"orig_iata"`
	DestIATA        string  `json:"destIata" db:"dest_iata"`
	DepartureTime   string  `json:"departureTime" db:"departure_time"`
	ArrivalTime     string  `json:"arrivalTime" db:"arrival_time"`
	TotalCash       float64 `json:"totalCash" db:"total_cash"`
	CorrectlyParsed bool    `json:"correctlyParsed" db:"correctly_parsed"`
}

func (view *FlightView) ToEntity(id oid.Id) *Flight {
	return &Flight{
		Id:              id,
		AirlCode:        view.AirlCode,
		FltNum:          view.FltNum,
		FltDate:         view.FltDate,
		OrigIATA:        view.OrigIATA,
		DestIATA:        view.DestIATA,
		DepartureTime:   view.DepartureTime,
		ArrivalTime:     view.ArrivalTime,
		TotalCash:       view.TotalCash,
		CorrectlyParsed: view.CorrectlyParsed,
	}
}
