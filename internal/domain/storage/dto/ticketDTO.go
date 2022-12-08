package dto

import (
	"github.com/achillescres/saina-api/internal/domain/entity"
	"github.com/achillescres/saina-api/pkg/object/oid"
)

type TicketCreate struct {
	Create          `json:"-" db:"-" binding:"-"`
	FlightId        oid.Id  `json:"flightId" db:"flight_id" binding:"required"`
	FlightNum       string  `json:"flightNum" db:"flight_num" binding:"required"`
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

func NewTicketCreate(
	flightId oid.Id,
	flightNum string,
	airlCode string,
	fltNum string,
	fltDate string,
	ticketCode string,
	ticketCapacity int,
	ticketType string,
	amount int,
	totalCash float64,
	correctlyParsed bool,
) *TicketCreate {
	return &TicketCreate{
		FlightId:        flightId,
		FlightNum:       flightNum,
		AirlCode:        airlCode,
		FltNum:          fltNum,
		FltDate:         fltDate,
		TicketCode:      ticketCode,
		TicketCapacity:  ticketCapacity,
		TicketType:      ticketType,
		Amount:          amount,
		TotalCash:       totalCash,
		CorrectlyParsed: correctlyParsed,
	}
}

func (tC *TicketCreate) ToEntity(id oid.Id) *entity.Ticket {
	return entity.NewTicket(
		id,
		tC.FlightId,
		tC.FlightNum,
		tC.AirlCode,
		tC.FltNum,
		tC.FltDate,
		tC.TicketCode,
		tC.TicketCapacity,
		tC.TicketType,
		tC.Amount,
		tC.TotalCash,
		tC.CorrectlyParsed,
	)
}
