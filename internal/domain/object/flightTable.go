package object

import "api-app/internal/domain/entity"

type FlightTable struct {
	entity.Flight
	Tickets []entity.Ticket
}

func NewFlightTable(flight entity.Flight, capacity int) *FlightTable {
	return &FlightTable{Flight: flight, Tickets: make([]entity.Ticket, 0, capacity)} //TODO gconfig injection
}

type FlightTableView struct {
	entity.FlightView
	TicketViews []entity.TicketView
}

func NewFlightTableView(flightView entity.FlightView, ticketViews []entity.TicketView) *FlightTableView {
	return &FlightTableView{FlightView: flightView, TicketViews: ticketViews}
}
