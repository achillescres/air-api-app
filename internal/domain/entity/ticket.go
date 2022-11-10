package entity

type TicketView struct {
	FlightId string `json:"flightId"`

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

type Ticket struct {
	Id   string
	View TicketView
}

func ToTicketView(t Ticket) TicketView {
	return t.View
}

func FromTicketView(id string, tV TicketView) *Ticket {
	return &Ticket{
		Id:   id,
		View: tV,
	}
}
