package entity

import "api-app/pkg/object/oid"

type Ticket struct {
	Entity `json:"-" db:"-" binding:"-"`
	Id     oid.Id `json:"id" db:"id" binding:"required"`

	// View
	FlightId        oid.Id  `json:"flightId" db:"flight_id" binding:"required"`
	AirlCode        string  `json:"airlCode" db:"airl_code" binding:"required"`
	FltNum          string  `json:"fltNum" db:"flt_num" binding:"required"`
	FltDate         string  `json:"fltDate" db:"flt_date" binding:"required"`
	TicketCode      string  `json:"ticketCode" db:"ticket_code" binding:"required"`
	TicketCapacity  int     `json:"ticketCapacity" db:"ticket_capacity" binding:"required"`
	TicketType      string  `json:"ticketType" db:"ticket_type" binding:"required"`
	Amount          int     `json:"amount" db:"amount" binding:"required"`
	TotalCash       float64 `json:"totalCash" db:"total_cash" binding:"required"`
	CorrectlyParsed bool    `json:"correct" db:"correctly_parsed" binding:"required"`
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
	View `json:"-" db:"-" binding:"-"`

	FlightId        oid.Id  `json:"flightId" db:"flight_id" binding:"required"`
	AirlCode        string  `json:"airlCode" db:"airl_code" binding:"required"`
	FltNum          string  `json:"fltNum" db:"flt_num" binding:"required"`
	FltDate         string  `json:"fltDate" db:"flt_date" binding:"required"`
	TicketCode      string  `json:"ticketCode" db:"ticket_code" binding:"required"`
	TicketCapacity  int     `json:"ticketCapacity" db:"ticket_capacity" binding:"required"`
	TicketType      string  `json:"ticketType" db:"ticket_type" binding:"required"`
	Amount          int     `json:"amount" db:"amount" binding:"required"`
	TotalCash       float64 `json:"totalCash" db:"total_cash" binding:"required"`
	CorrectlyParsed bool    `json:"correct" db:"correctly_parsed" binding:"required"`
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
