package entity

type Ticket struct {
	Id              string  `json:"id"`
	FlightId        string  `json:"flightId"`
	AirlCode        string  `json:"airlCode"`
	FltNum          string  `json:"fltNum"`
	FltDate         string  `json:"fltDate"`
	TicketCode      string  `json:"ticketCode"`
	TicketType      string  `json:"ticketType"`
	Amount          int     `json:"amount"`
	TotalCash       float64 `json:"totalCash"`
	CorrectlyParsed bool    `json:"correct"`
}

type TicketView struct {
	FlightId        string  `json:"flightId"`
	AirlCode        string  `json:"airlCode"`
	FltNum          string  `json:"fltNum"`
	FltDate         string  `json:"fltDate"`
	TicketCode      string  `json:"ticketCode"`
	TicketType      string  `json:"ticketType"`
	Amount          int     `json:"amount"`
	TotalCash       float64 `json:"totalCash"`
	CorrectlyParsed bool    `json:"correct"`
}

func ToTicketView(t Ticket) *TicketView {
	return &TicketView{
		FlightId:        t.FlightId,
		AirlCode:        t.AirlCode,
		FltNum:          t.FltNum,
		FltDate:         t.FltDate,
		TicketCode:      t.TicketCode,
		TicketType:      t.TicketType,
		Amount:          t.Amount,
		TotalCash:       t.TotalCash,
		CorrectlyParsed: t.CorrectlyParsed,
	}
}

func FromTicketView(id string, tV TicketView) *Ticket {
	return &Ticket{
		Id:              id,
		FlightId:        tV.FlightId,
		AirlCode:        tV.AirlCode,
		FltNum:          tV.FltNum,
		FltDate:         tV.FltDate,
		TicketCode:      tV.TicketCode,
		TicketType:      tV.TicketType,
		Amount:          tV.Amount,
		TotalCash:       tV.TotalCash,
		CorrectlyParsed: tV.CorrectlyParsed,
	}
}
