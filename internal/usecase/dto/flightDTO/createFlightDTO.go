package flightDTO

import (
	"api-app/internal/domain/entity"
)

type CreateFlightDTO struct {
	AirlCode        string  `json:"airlCode"`
	FltNum          string  `json:"fltNum"`
	FltDate         string  `json:"fltDate"`
	OrigIATA        string  `json:"origIata"`
	DestIATA        string  `json:"destIata"`
	TotalCash       float64 `json:"totalCash"`
	CorrectlyParsed bool    `json:"correctlyParsed"`
}

func NewCreateFlightDTO(
	airlCode string,
	fltNum string,
	fltDate string,
	origIATA string,
	destIATA string,
	totalCash float64,
	correctlyParsed bool,
) *CreateFlightDTO {
	return &CreateFlightDTO{
		AirlCode:        airlCode,
		FltNum:          fltNum,
		FltDate:         fltDate,
		OrigIATA:        origIATA,
		DestIATA:        destIATA,
		TotalCash:       totalCash,
		CorrectlyParsed: correctlyParsed}
}

func NewCreateFlightDTOFromFlight(f entity.Flight) *CreateFlightDTO {
	return &CreateFlightDTO{
		AirlCode:        f.AirlCode,
		FltNum:          f.FltNum,
		FltDate:         f.FltDate,
		OrigIATA:        f.OrigIATA,
		DestIATA:        f.DestIATA,
		TotalCash:       f.TotalCash,
		CorrectlyParsed: f.CorrectlyParsed,
	}
}

func NewFlightFromCreateFlightDTO(id string, cFDTO CreateFlightDTO) *entity.Flight {
	return entity.NewFlight(
		id,
		cFDTO.AirlCode,
		cFDTO.FltNum,
		cFDTO.OrigIATA,
		cFDTO.DestIATA,
		cFDTO.TotalCash,
		cFDTO.FltDate,
		cFDTO.CorrectlyParsed,
	)
}
