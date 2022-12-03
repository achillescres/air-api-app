package object

import (
	entity2 "github.com/achillescres/saina-api/internal/domain/entity"
)

type FlightTable struct {
	entity2.Flight
	Tickets []entity2.Ticket `json:"tickets" binding:"required"`
}

func NewFlightTable(flight entity2.Flight, capacity int) *FlightTable {
	return &FlightTable{Flight: flight, Tickets: make([]entity2.Ticket, 0, capacity)} //TODO gconfig injection
}
