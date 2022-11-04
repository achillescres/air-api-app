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

type FlightView struct {
	AirlCode        string  `json:"airlCode"`
	FltNum          string  `json:"fltNum"`
	FltDate         string  `json:"fltDate"`
	OrigIATA        string  `json:"origIata"`
	DestIATA        string  `json:"destIata"`
	TotalCash       float64 `json:"totalCash"`
	CorrectlyParsed bool    `json:"correctlyParsed"`
}

func FromFlightView(id string, view FlightView) *Flight {
	return &Flight{
		Id:              id,
		AirlCode:        view.AirlCode,
		FltNum:          view.FltNum,
		FltDate:         view.FltDate,
		OrigIATA:        view.OrigIATA,
		DestIATA:        view.DestIATA,
		TotalCash:       view.TotalCash,
		CorrectlyParsed: view.CorrectlyParsed}
}

func ToFlightView(id string, view FlightView) *Flight {
	return &Flight{
		Id:              id,
		AirlCode:        view.AirlCode,
		FltNum:          view.FltNum,
		FltDate:         view.FltNum,
		OrigIATA:        view.OrigIATA,
		DestIATA:        view.DestIATA,
		TotalCash:       view.TotalCash,
		CorrectlyParsed: view.CorrectlyParsed,
	}
}
