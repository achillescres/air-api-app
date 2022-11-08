package entity

import "time"

type Flight struct {
	Id string `json:"id"`

	AirlCode string `json:"airlCode"`
	FltNum   string `json:"fltNum"`
	FltDate  string `json:"fltDate"`

	OrigIATA string `json:"origIata"`
	DestIATA string `json:"destIata"`

	DepartureTime time.Time `json:"departureTime"`
	ArriveTime    time.Time `json:"arriveTime"`

	TotalCash float64 `json:"totalCash"`

	CorrectlyParsed bool `json:"correctlyParsed"`
}

type FlightView struct {
	AirlCode string `json:"airlCode"`

	FltNum  string `json:"fltNum"`
	FltDate string `json:"fltDate"`

	OrigIATA string `json:"origIata"`
	DestIATA string `json:"destIata"`

	DepartureTime time.Time `json:"departureTime"`
	ArriveTime    time.Time `json:"arriveTime"`

	TotalCash float64 `json:"totalCash"`

	CorrectlyParsed bool `json:"correctlyParsed"`
}

func FromFlightView(id string, view FlightView) *Flight {
	return &Flight{
		Id: id,

		AirlCode: view.AirlCode,

		FltNum:  view.FltNum,
		FltDate: view.FltDate,

		OrigIATA: view.OrigIATA,
		DestIATA: view.DestIATA,

		DepartureTime: view.DepartureTime,
		ArriveTime:    view.ArriveTime,

		TotalCash:       view.TotalCash,
		CorrectlyParsed: view.CorrectlyParsed}
}

func ToFlightView(id string, f Flight) *FlightView {
	return &FlightView{
		AirlCode: f.AirlCode,

		FltNum:  f.FltNum,
		FltDate: f.FltNum,

		OrigIATA: f.OrigIATA,
		DestIATA: f.DestIATA,

		DepartureTime: f.DepartureTime,
		ArriveTime:    f.ArriveTime,

		TotalCash: f.TotalCash,

		CorrectlyParsed: f.CorrectlyParsed,
	}
}
