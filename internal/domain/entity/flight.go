package entity

import (
	"api-app/pkg/object/oid"
)

type Flight struct {
	Entity `json:"-" db:"-" binding:"-"`
	Id     oid.Id `json:"id" db:"id" binding:"required"`

	// View
	AirlCode        string  `json:"airlCode" db:"airl_code" binding:"required"`
	FltNum          string  `json:"fltNum" db:"flt_num" binding:"required"`
	FltDate         string  `json:"fltDate" db:"flt_date" binding:"required"`
	OrigIATA        string  `json:"origIata" db:"orig_iata" binding:"required"`
	DestIATA        string  `json:"destIata" db:"dest_iata" binding:"required"`
	DepartureTime   string  `json:"departureTime" db:"departure_time" binding:"required"`
	ArrivalTime     string  `json:"arrivalTime" db:"arrival_time" binding:"required"`
	TotalCash       float64 `json:"totalCash" db:"total_cash" binding:"required"`
	CorrectlyParsed bool    `json:"correctlyParsed" db:"correctly_parsed" binding:"required"`
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
	AirlCode        string  `json:"airlCode" db:"airl_code" binding:"required"`
	FltNum          string  `json:"fltNum" db:"flt_num" binding:"required"`
	FltDate         string  `json:"fltDate" db:"flt_date" binding:"required"`
	OrigIATA        string  `json:"origIata" db:"orig_iata" binding:"required"`
	DestIATA        string  `json:"destIata" db:"dest_iata" binding:"required"`
	DepartureTime   string  `json:"departureTime" db:"departure_time" binding:"required"`
	ArrivalTime     string  `json:"arrivalTime" db:"arrival_time" binding:"required"`
	TotalCash       float64 `json:"totalCash" db:"total_cash" binding:"required"`
	CorrectlyParsed bool    `json:"correctlyParsed" db:"correctly_parsed" binding:"required"`
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
