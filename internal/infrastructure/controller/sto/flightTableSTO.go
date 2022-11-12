package sto

import (
	"api-app/internal/domain/object"
)

type FlightTableSTO struct {
	Flight  FlightSTO
	Tickets []TicketSTO
}

func ToFLightTableSTO(fT object.FlightTable) *FlightTableSTO {
	ticketSTos := make([]TicketSTO, 0, len(fT.Tickets))
	for _, ticket := range fT.Tickets {
		ticketSTos = append(ticketSTos, ToTicketSTO(ticket))
	}

	return &FlightTableSTO{
		Flight:  ToFlightSTO(fT.Flight),
		Tickets: ticketSTos,
	}
}

type FlightTableViewSTO struct {
	FlightView  FlightViewSTO
	TicketViews []TicketViewSTO
}
