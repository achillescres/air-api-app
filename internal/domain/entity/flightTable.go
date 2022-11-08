package entity

type FlightTable struct {
	Flight
	Tickets []Ticket
}

func NewFlightTable(flight Flight, capacity int) *FlightTable {
	return &FlightTable{Flight: flight, Tickets: make([]Ticket, 0, capacity)} //TODO gconfig injection
}

type FlightTableView struct {
	FlightView
	TicketViews []TicketView
}

func NewFlightTableView(flightView FlightView, ticketViews []TicketView) *FlightTableView {
	return &FlightTableView{FlightView: flightView, TicketViews: ticketViews}
}
