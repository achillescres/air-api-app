package dto

import (
	"github.com/achillescres/saina-api/internal/domain/entity"
	"github.com/achillescres/saina-api/pkg/object/oid"
)

type TicketCreate struct {
	Create
	FlightId        oid.Id  `json:"flightId" `
	AirlCode        string  `json:"airlCode" `
	FltNum          string  `json:"fltNum" `
	FltDate         string  `json:"fltDate" `
	TicketCode      string  `json:"ticketCode" `
	TicketCapacity  int     `json:"ticketCapacity" `
	TicketType      string  `json:"ticketType" `
	Amount          int     `json:"amount" `
	TotalCash       float64 `json:"totalCash" `
	CorrectlyParsed bool    `json:"correct" `
}

func (tC *TicketCreate) ToEntity(id oid.Id) *entity.Ticket {
	return &entity.Ticket{
		Id:              id,
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
