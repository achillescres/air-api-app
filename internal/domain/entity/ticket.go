package entity

import (
	"github.com/achillescres/saina-api/pkg/object/oid"
)

type Ticket struct {
	Entity          `json:"-" db:"-" binding:"-"`
	Id              oid.Id  `json:"id" db:"id" binding:"required"`
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
