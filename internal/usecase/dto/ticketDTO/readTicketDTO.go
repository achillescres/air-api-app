package ticketDTO

type ReadTicketDTO struct {
	Id         string  `json:"id"`
	FlightId   string  `json:"flightId"`
	AirlCode   string  `json:"airlCode"`
	FltNum     string  `json:"fltNum"`
	FltDate    string  `json:"fltDate"`
	TicketCode string  `json:"ticketCode"`
	TicketType string  `json:"ticketType"`
	Amount     int     `json:"amount"`
	TotalCash  float64 `json:"totalCash"`
	Correct    bool    `json:"correct"`
}

func NewReadTicketDTO(
	id string,
	flightId string,
	airlCode string,
	fltNum string,
	fltDate string,
	ticketCode string,
	ticketType string,
	amount int,
	totalCash float64,
	correct bool,
) *ReadTicketDTO {
	return &ReadTicketDTO{Id: id,
		FlightId:   flightId,
		AirlCode:   airlCode,
		FltNum:     fltNum,
		FltDate:    fltDate,
		TicketCode: ticketCode,
		TicketType: ticketType,
		Amount:     amount,
		TotalCash:  totalCash,
		Correct:    correct}
}