package product

import (
	httpHandler "api-app/internal/infrastructure/controller/handler/http"
	"context"
	"github.com/gin-gonic/gin"
)

type Handlers interface {
	RegisterAll(ctx context.Context, r *gin.RouterGroup) error
}

type handlers struct {
	flightHand httpHandler.FlightHandler
	ticketHand httpHandler.TicketHandler
	userHand   httpHandler.UserHandler
}

func NewHandlers(usecase Usecases) (Handlers, error) {
	return &handlers{
		flightHand: httpHandler.NewFlightHandler(usecase.FlightUc()),
		ticketHand: httpHandler.NewTicketHandler(usecase.TicketUc()),
		userHand:   httpHandler.NewUserHandler(usecase.UserUc()),
	}, nil
}

func (h *handlers) RegisterAll(ctx context.Context, r *gin.RouterGroup) error {
	api := r.Group("/api")
	{
		h.userHand.RegisterRouter(api)
		h.ticketHand.RegisterRouter(api)
		h.flightHand.RegisterRouter(api)
	}

	return nil
}
