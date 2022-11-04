package sto

import "api-app/internal/domain/entity"

type FlightSTO entity.Flight

func ToFlightSTO(f entity.Flight) FlightSTO {
	return FlightSTO(f)
}

type FlightViewSTO entity.FlightView

func ToFlightViewSTO(f entity.FlightView) FlightViewSTO {
	return FlightViewSTO(f)
}
