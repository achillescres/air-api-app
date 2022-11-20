package usecase

import (
	"api-app/internal/domain/entity"
	"api-app/internal/domain/service"
	"api-app/internal/domain/storage/dto"
	"api-app/pkg/object/oid"
	"context"
)

type TicketUsecase interface {
	Usecase
	StoreTicket(ctx context.Context, tC dto.TicketCreate) (*entity.Ticket, error)
	GetAllTicketsByMap(ctx context.Context) (map[oid.Id]*entity.Ticket, error)
}

type ticketUsecase struct {
	ticketService service.TicketService
}

var _ TicketUsecase = (*ticketUsecase)(nil)

func NewTicketUsecase(ticketService service.TicketService) TicketUsecase {
	return &ticketUsecase{ticketService}
}

func (tU *ticketUsecase) GetAllTicketsByMap(ctx context.Context) (map[oid.Id]*entity.Ticket, error) {
	byMap, err := tU.ticketService.GetAllByMap(ctx)
	if err != nil {
		return nil, err
	}
	return byMap, err
}

func (tU *ticketUsecase) StoreTicket(ctx context.Context, tC dto.TicketCreate) (*entity.Ticket, error) {
	store, err := tU.ticketService.Store(ctx, tC)
	if err != nil {
		return nil, err
	}
	return store, nil
}
