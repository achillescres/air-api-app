package dto

import (
	"api-app/internal/domain/entity"
	"api-app/pkg/object/oid"
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
