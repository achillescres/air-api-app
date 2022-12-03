package dto

import (
	"github.com/achillescres/saina-api/internal/domain/entity"
	"github.com/achillescres/saina-api/pkg/object/oid"
)

type FLightCreate struct {
	Create
	AirlCode        string  `json:"airlCode" `
	FltNum          string  `json:"fltNum" `
	FltDate         string  `json:"fltDate" `
	OrigIATA        string  `json:"origIata" `
	DestIATA        string  `json:"destIata" `
	DepartureTime   string  `json:"departureTime" `
	ArrivalTime     string  `json:"arrivalTime" `
	TotalCash       float64 `json:"totalCash" `
	CorrectlyParsed bool    `json:"correctlyParsed" `
}

func (fC *FLightCreate) ToEntity(id oid.Id) *entity.Flight {
	return &entity.Flight{
		Id:              id,
		AirlCode:        fC.AirlCode,
		FltNum:          fC.FltNum,
		FltDate:         fC.FltDate,
		OrigIATA:        fC.OrigIATA,
		DestIATA:        fC.DestIATA,
		DepartureTime:   fC.DepartureTime,
		ArrivalTime:     fC.ArrivalTime,
		TotalCash:       fC.TotalCash,
		CorrectlyParsed: fC.CorrectlyParsed,
	}
}
