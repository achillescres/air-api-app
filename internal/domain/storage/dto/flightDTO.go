package dto

import (
	"api-app/internal/domain/entity"
	"time"
)

type FLightCreate struct {
	Create
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

func (fC *FLightCreate) ToView() *entity.FlightView {
	return &entity.FlightView{
		AirlCode:        fC.AirlCode,
		FltNum:          fC.FltNum,
		FltDate:         fC.FltDate,
		OrigIATA:        fC.OrigIATA,
		DestIATA:        fC.DestIATA,
		DepartureTime:   fC.DepartureTime,
		ArriveTime:      fC.ArriveTime,
		TotalCash:       fC.TotalCash,
		CorrectlyParsed: fC.CorrectlyParsed,
	}
}
