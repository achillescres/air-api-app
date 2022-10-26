package ticketDTO

import (
	"api-app/internal/entity"
)

type CreateTicketDTO struct {
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

func NewCreateTicketDTO(
	flightId string,
	airlCode string,
	fltNum string,
	fltDate string,
	ticketCode string,
	ticketType string,
	amount int,
	totalCash float64,
	correct bool,
) *CreateTicketDTO {
	return &CreateTicketDTO{
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

func NewTicketFromCreateTicketDTO(
	id string,
	cTDTO CreateTicketDTO,
) *entity.Ticket {
	return entity.NewTicket(
		id,
		cTDTO.FlightId,
		cTDTO.AirlCode,
		cTDTO.FltNum,
		cTDTO.FltDate,
		cTDTO.TicketCode,
		cTDTO.TicketType,
		cTDTO.Amount,
		cTDTO.TotalCash,
		cTDTO.Correct)
}
