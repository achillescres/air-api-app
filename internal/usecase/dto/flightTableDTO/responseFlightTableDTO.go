package flightTableDTO

import "api-app/internal/entity"

type ResponseFlightTableDTO struct {
	Id              string           `json:"id"`
	AirlCode        string           `json:"airlCode"`
	FltNum          string           `json:"fltNum"`
	FltDate         string           `json:"fltDate"`
	OrigIATA        string           `json:"origIata"`
	DestIATA        string           `json:"destIata"`
	TotalCash       float64          `json:"totalCash"`
	CorrectlyParsed bool             `json:"correctlyParsed"`
	Tickets         []*entity.Ticket `json:"tickets"`
}

func NewResponseFlightTableDTO(
	id string,
	airlCode string,
	fltNum string,
	fltDate string,
	origIATA string,
	destIATA string,
	totalCash float64,
	correctlyParsed bool,
	tickets []*entity.Ticket,
) *ResponseFlightTableDTO {
	return &ResponseFlightTableDTO{
		Id:              id,
		AirlCode:        airlCode,
		FltNum:          fltNum,
		FltDate:         fltDate,
		OrigIATA:        origIATA,
		DestIATA:        destIATA,
		TotalCash:       totalCash,
		CorrectlyParsed: correctlyParsed,
		Tickets:         tickets,
	}
}

func NewResponseFlightTableDTOFromFlight(f entity.Flight, tickets []*entity.Ticket) *ResponseFlightTableDTO {
	return &ResponseFlightTableDTO{
		Id:              f.Id,
		AirlCode:        f.AirlCode,
		FltNum:          f.FltNum,
		FltDate:         f.FltDate,
		OrigIATA:        f.OrigIATA,
		DestIATA:        f.DestIATA,
		TotalCash:       f.TotalCash,
		CorrectlyParsed: f.CorrectlyParsed,
		Tickets:         tickets,
	}
}
