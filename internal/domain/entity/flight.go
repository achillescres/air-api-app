package entity

import (
	"github.com/achillescres/saina-api/pkg/object/oid"
)

type Flight struct {
	Entity          `json:"-" db:"-" binding:"-"`
	Id              oid.Id  `json:"id" db:"id" binding:"required"`
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

func NewFlight(
	id oid.Id,
	airlCode string,
	fltNum string,
	fltDate string,
	origIATA string,
	destIATA string,
	departureTime string,
	arrivalTime string,
	totalCash float64,
	correctlyParsed bool,
) *Flight {
	return &Flight{
		Id:              id,
		AirlCode:        airlCode,
		FltNum:          fltNum,
		FltDate:         fltDate,
		OrigIATA:        origIATA,
		DestIATA:        destIATA,
		DepartureTime:   departureTime,
		ArrivalTime:     arrivalTime,
		TotalCash:       totalCash,
		CorrectlyParsed: correctlyParsed,
	}
}
