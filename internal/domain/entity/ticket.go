package entity

import "api-app/pkg/object/oid"

type Ticket struct {
	Entity
	Id   oid.Id     `json:"id" binding:"required"`
	View TicketView `json:"view" binding:"required"`
}

type TicketView struct {
	View
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

func ToTicketView(t Ticket) TicketView {
	return t.View
}

func FromTicketView(id oid.Id, tV TicketView) *Ticket {
	return &Ticket{
		Id:   id,
		View: tV,
	}
}
