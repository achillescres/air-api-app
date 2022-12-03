package sto

import (
	"github.com/achillescres/saina-api/internal/domain/object"
)

type FlightTableSTO struct {
	object.FlightTable
}

func ToFlightTableSTO(fT *object.FlightTable) *FlightTableSTO {
	return &FlightTableSTO{
		FlightTable: *fT,
	}
}
