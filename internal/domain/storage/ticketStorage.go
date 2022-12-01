package storage

import (
	"api-app/internal/domain/entity"
	"api-app/internal/domain/storage/dto"
	"api-app/pkg/object/oid"
	"context"
)

type TicketStorage interface {
	Storage[entity.Ticket, dto.TicketCreate]
	GetAllByFlightId(ctx context.Context, flightId oid.Id) ([]*entity.Ticket, error)
}
