package entity

type Flight struct {
	Id              string  `json:"id"`
	AirlCode        string  `json:"airlCode"`
	FltNum          string  `json:"fltNum"`
	FltDate         string  `json:"fltDate"`
	OrigIATA        string  `json:"origIata"`
	DestIATA        string  `json:"destIata"`
	TotalCash       float64 `json:"totalCash"`
	CorrectlyParsed bool    `json:"correctlyParsed"`
}

func NewFlight(
	id string,
	airlCode string,
	fltNum string,
	origIATA string,
	destIATA string,
	totalCash float64,
	fltDate string,
	correct bool,
) *Flight {
	return &Flight{
		Id:              id,
		AirlCode:        airlCode,
		FltNum:          fltNum,
		FltDate:         fltDate,
		OrigIATA:        origIATA,
		DestIATA:        destIATA,
		TotalCash:       totalCash,
		CorrectlyParsed: correct,
	}
}
