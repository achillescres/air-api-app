package storage

import (
	"context"
	"github.com/achillescres/saina-api/internal/domain/entity"
	"github.com/achillescres/saina-api/internal/domain/storage/dto"
	"github.com/achillescres/saina-api/pkg/object/oid"
)

type TicketStorage interface {
	Storage[entity.Ticket, dto.TicketCreate]
	GetAllByFlightId(ctx context.Context, flightId oid.Id) ([]*entity.Ticket, error)
}
