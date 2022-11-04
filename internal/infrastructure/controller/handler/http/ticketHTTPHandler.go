package httpHandler

import (
	"api-app/internal/domain/usecase"
)

type TicketHandler interface {
	usecase.TicketUsecase
}

var _ Handler = (*TicketHandler)(nil)

type ticketHandler struct {
	usecase.TicketUsecase
}

var _ TicketHandler = (*ticketHandler)(nil)

func NewTicketHandler(tUc usecase.TicketUsecase) *ticketHandler {
	return &ticketHandler{TicketUsecase: tUc}
}
