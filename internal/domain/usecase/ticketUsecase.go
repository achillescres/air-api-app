package usecase

import (
	"api-app/internal/domain/service"
)

type TicketUsecase interface {
	service.TicketService
}

type ticketUsecase struct {
	service.TicketService
}

func NewTicketUsecase(ticketService service.TicketService) *ticketUsecase {
	return &ticketUsecase{ticketService}
}
