package storage

import (
	"api-app/internal/domain/entity"
)

type TicketStorage interface {
	GetById(id string) entity.Ticket
	GetAll() []entity.Ticket
	GetAllByFlightId(flightId string) []entity.Ticket

	Store(f entity.Ticket) error
	DeleteById(id string) (entity.Ticket, error)
}

var _ Storage = (*TicketStorage)(nil)
