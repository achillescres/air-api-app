package entity

import "api-app/pkg/object/oid"

type Ticket struct {
	Id   oid.Id `json:"id"`
	View TicketView
}

var _ Entity = (*Ticket)(nil)

type TicketView struct {
	FlightId oid.Id `json:"flightId"`

	AirlCode string `json:"airlCode"`

	FltNum  string `json:"fltNum"`
	FltDate string `json:"fltDate"`

	TicketCode     string `json:"ticketCode"`
	TicketCapacity int    `json:"ticketCapacity"`
	TicketType     string `json:"ticketType"`

	Amount int `json:"amount"`

	TotalCash float64 `json:"totalCash"`

	CorrectlyParsed bool `json:"correct"`
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
