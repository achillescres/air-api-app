package safeObject

import (
	"github.com/achillescres/saina-api/internal/domain/object"
)

type FlightTableSafe struct {
	object.FlightTable
}

func ToFlightTableSafe(fT *object.FlightTable) *FlightTableSafe {
	return &FlightTableSafe{
		FlightTable: *fT,
	}
}
