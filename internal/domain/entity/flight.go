package entity

import (
	"api-app/pkg/object/oid"
	"time"
)

type Flight struct {
	Entity
	Id   oid.Id     `json:"id" binding:"required"`
	View FlightView `json:"view" binding:"required"`
}

type FlightView struct {
	View
	AirlCode        string    `json:"airlCode" binding:"required"`
	FltNum          string    `json:"fltNum" binding:"required"`
	FltDate         string    `json:"fltDate" binding:"required"`
	OrigIATA        string    `json:"origIata" binding:"required"`
	DestIATA        string    `json:"destIata" binding:"required"`
	DepartureTime   time.Time `json:"departureTime" binding:"required"`
	ArriveTime      time.Time `json:"arriveTime" binding:"required"`
	TotalCash       float64   `json:"totalCash" binding:"required"`
	CorrectlyParsed bool      `json:"correctlyParsed" binding:"required"`
}

func NewFlightView(airlCode string, fltNum string, fltDate string, origIATA string, destIATA string, departureTime time.Time, arriveTime time.Time, totalCash float64, correctlyParsed bool) *FlightView {
	return &FlightView{AirlCode: airlCode, FltNum: fltNum, FltDate: fltDate, OrigIATA: origIATA, DestIATA: destIATA, DepartureTime: departureTime, ArriveTime: arriveTime, TotalCash: totalCash, CorrectlyParsed: correctlyParsed}
}

func FromFlightView(id oid.Id, view FlightView) *Flight {
	return &Flight{
		Id:   id,
		View: view,
	}
}

func ToFlightView(f Flight) FlightView {
	return f.View
}
