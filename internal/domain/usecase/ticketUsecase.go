package usecase

import (
	"api-app/internal/domain/entity"
	"api-app/internal/domain/service"
)

type TicketUsecase interface {
	Usecase[entity.Ticket, entity.TicketView]
}

type ticketUsecase struct {
	service.TicketService
}

var _ TicketUsecase = (*ticketUsecase)(nil)

func NewTicketUsecase(ticketService service.TicketService) TicketUsecase {
	return &ticketUsecase{ticketService}
}
