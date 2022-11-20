package entity

import "api-app/pkg/object/oid"

type Ticket struct {
	Entity `json:"-" db:"-"`
	Id     oid.Id `json:"id" db:"id"`

	// View
	FlightId        oid.Id  `json:"flightId" db:"flight_id"`
	AirlCode        string  `json:"airlCode" db:"airl_code"`
	FltNum          string  `json:"fltNum" db:"flt_num"`
	FltDate         string  `json:"fltDate" db:"flt_date"`
	TicketCode      string  `json:"ticketCode" db:"ticket_code"`
	TicketCapacity  int     `json:"ticketCapacity" db:"ticket_capacity"`
	TicketType      string  `json:"ticketType" db:"ticket_type"`
	Amount          int     `json:"amount" db:"amount"`
	TotalCash       float64 `json:"totalCash" db:"total_cash"`
	CorrectlyParsed bool    `json:"correct" db:"correctly_parsed"`
}

var _ Entity = (*Ticket)(nil)

func (t *Ticket) ToView() *TicketView {
	return &TicketView{
		FlightId:        t.FlightId,
		AirlCode:        t.AirlCode,
		FltNum:          t.FltNum,
		FltDate:         t.FltDate,
		TicketCode:      t.TicketCode,
		TicketCapacity:  t.TicketCapacity,
		TicketType:      t.TicketType,
		Amount:          t.Amount,
		TotalCash:       t.TotalCash,
		CorrectlyParsed: t.CorrectlyParsed,
	}
}

type TicketView struct {
	View            `json:"-" db:"-"`
	FlightId        oid.Id  `json:"flightId" db:"flight_id"`
	AirlCode        string  `json:"airlCode" db:"airl_code"`
	FltNum          string  `json:"fltNum" db:"flt_num"`
	FltDate         string  `json:"fltDate" db:"flt_date"`
	TicketCode      string  `json:"ticketCode" db:"ticket_code"`
	TicketCapacity  int     `json:"ticketCapacity" db:"ticket_capacity"`
	TicketType      string  `json:"ticketType" db:"ticket_type"`
	Amount          int     `json:"amount" db:"amount"`
	TotalCash       float64 `json:"totalCash" db:"total_cash"`
	CorrectlyParsed bool    `json:"correct" db:"correctly_parsed"`
}

var _ View = (*TicketView)(nil)

func (tV *TicketView) ToEntity(id oid.Id) *Ticket {
	return &Ticket{
		Id:              id,
		FlightId:        tV.FlightId,
		AirlCode:        tV.AirlCode,
		FltNum:          tV.FltNum,
		FltDate:         tV.FltDate,
		TicketCode:      tV.TicketCode,
		TicketCapacity:  tV.TicketCapacity,
		TicketType:      tV.TicketType,
		Amount:          tV.Amount,
		TotalCash:       tV.TotalCash,
		CorrectlyParsed: tV.CorrectlyParsed,
	}
}
