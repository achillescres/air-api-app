package service

import (
	"context"
	"github.com/achillescres/saina-api/internal/domain/entity"
	"github.com/achillescres/saina-api/internal/domain/storage"
	"github.com/achillescres/saina-api/internal/domain/storage/dto"
	"github.com/achillescres/saina-api/pkg/object/oid"
)

type TicketService PrimitiveService[entity.Ticket, dto.TicketCreate]

type ticketService struct {
	TicketService
	storage storage.TicketStorage
}

var _ TicketService = (*ticketService)(nil)

func NewTicketService(storage storage.TicketStorage) TicketService {
	return &ticketService{storage: storage}
}

func (tService *ticketService) GetAllByMap(ctx context.Context) (map[oid.Id]*entity.Ticket, error) {
	ticketsMap := map[oid.Id]*entity.Ticket{}
	tickets, err := tService.storage.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	for _, ticket := range tickets {
		ticketsMap[ticket.Id] = ticket
	}

	return ticketsMap, nil
}

func (tService *ticketService) GetById(ctx context.Context, id oid.Id) (*entity.Ticket, error) {
	return tService.GetById(ctx, id)
}

func (tService *ticketService) GetAll(ctx context.Context) ([]*entity.Ticket, error) {
	return tService.storage.GetAll(ctx)
}

func (tService *ticketService) Store(ctx context.Context, tC dto.TicketCreate) (*entity.Ticket, error) {
	return tService.storage.Store(ctx, tC)
}

func (tService *ticketService) DeleteById(ctx context.Context, id oid.Id) (*entity.Ticket, error) {
	return tService.storage.DeleteById(ctx, id)
}
