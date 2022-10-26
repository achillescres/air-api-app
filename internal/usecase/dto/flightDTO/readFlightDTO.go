package flightDTO

type ReadFlightDTO struct {
	Id              string  `json:"id"`
	AirlCode        string  `json:"airlCode"`
	FltNum          string  `json:"fltNum"`
	FltDate         string  `json:"fltDate"`
	OrigIATA        string  `json:"origIata"`
	DestIATA        string  `json:"destIata"`
	TotalCash       float64 `json:"totalCash"`
	CorrectlyParsed bool    `json:"correctlyParsed"`
}

func NewReadFlightDTO(
	id string,
	airlCode string,
	fltNum string,
	fltDate string,
	origIATA string,
	destIATA string,
	totalCash float64,
	correctlyParsed bool,
) *ReadFlightDTO {
	return &ReadFlightDTO{
		Id:              id,
		AirlCode:        airlCode,
		FltNum:          fltNum,
		FltDate:         fltDate,
		OrigIATA:        origIATA,
		DestIATA:        destIATA,
		TotalCash:       totalCash,
		CorrectlyParsed: correctlyParsed}
}
