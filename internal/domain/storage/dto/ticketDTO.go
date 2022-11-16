package dto

import (
	"api-app/internal/domain/entity"
	"api-app/pkg/object/oid"
)

type TicketCreate struct {
	Create
	FlightId        oid.Id  `json:"flightId" binding:"required"`
	AirlCode        string  `json:"airlCode" binding:"required"`
	FltNum          string  `json:"fltNum" binding:"required"`
	FltDate         string  `json:"fltDate" binding:"required"`
	TicketCode      string  `json:"ticketCode" binding:"required"`
	TicketCapacity  int     `json:"ticketCapacity" binding:"required"`
	TicketType      string  `json:"ticketType" binding:"required"`
	Amount          int     `json:"amount" binding:"required"`
	TotalCash       float64 `json:"totalCash" binding:"required"`
	CorrectlyParsed bool    `json:"correct" binding:"required"`
}

func (tC *TicketCreate) ToTicketView() *entity.TicketView {
	return &entity.TicketView{
		FlightId:        tC.FlightId,
		AirlCode:        tC.AirlCode,
		FltNum:          tC.FltNum,
		FltDate:         tC.FltDate,
		TicketCode:      tC.TicketCode,
		TicketCapacity:  tC.TicketCapacity,
		TicketType:      tC.TicketType,
		Amount:          tC.Amount,
		TotalCash:       tC.TotalCash,
		CorrectlyParsed: tC.CorrectlyParsed,
	}
}
