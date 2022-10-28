package handler

import (
	"api-app/internal/usecase"
)

type TicketHandler interface {
}

var _ Handler = (*TicketHandler)(nil)

type ticketHandler struct {
	uc usecase.TicketUsecase
}

var _ TicketHandler = (*ticketHandler)(nil)

func NewTicketHandler(uc usecase.TicketUsecase) *ticketHandler {
	return &ticketHandler{uc: uc}
}
