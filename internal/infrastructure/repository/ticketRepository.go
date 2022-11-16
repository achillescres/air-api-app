package repository

import (
	"api-app/internal/config"
	"api-app/internal/domain/entity"
	"api-app/internal/domain/storage"
	"api-app/internal/domain/storage/dto"
	"api-app/pkg/db/postgresql"
	"api-app/pkg/object/oid"
	"context"
)

type TicketRepository interface {
	storage.Storage[entity.Ticket, entity.TicketView, dto.TicketCreate]
}

type ticketRepository struct {
	pool postgresql.Pool
	cfg  config.DBConfig
}

var _ storage.TicketStorage = (*ticketRepository)(nil)

func NewTicketRepository(pool postgresql.Pool, cfg config.DBConfig) TicketRepository {
	return &ticketRepository{pool: pool, cfg: cfg}
}

func (tRepo *ticketRepository) GetById(ctx context.Context, id oid.Id) (entity.Ticket, error) {
	// TODO implement me
	panic("implement me")
}

func (tRepo *ticketRepository) GetAll(ctx context.Context) ([]entity.Ticket, error) {
	// TODO implement me
	panic("implement me")
}

func (tRepo *ticketRepository) GetAllByFlightId(ctx context.Context, flightId oid.Id) ([]entity.Ticket, error) {
	// TODO implement me
	panic("implement me")
}

func (tRepo *ticketRepository) Store(ctx context.Context, tC dto.TicketCreate) (entity.Ticket, error) {
	// TODO impl me
	panic("implement me")
}

func (tRepo *ticketRepository) DeleteById(ctx context.Context, id oid.Id) (entity.Ticket, error) {
	// TODO implement me
	panic("implement me")
}
