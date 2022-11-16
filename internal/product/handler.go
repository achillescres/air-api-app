package product

import (
	httpHandler "api-app/internal/infrastructure/controller/handler/http"
	"context"
	"github.com/gin-gonic/gin"
)

type Handler interface {
	//FlightHand() httpHandler.FlightHandler
	//TickerHand() httpHandler.TicketHandler
	//UserHand() httpHandler.UserHandler
	RegisterAll(ctx context.Context, r *gin.RouterGroup) error
}

type handler struct {
	flightHand httpHandler.FlightHandler
	ticketHand httpHandler.TicketHandler
	userHand   httpHandler.UserHandler
}

//func (h *handler) FlightHand() httpHandler.FlightHandler {
//	return h.flightHand
//}
//
//func (h *handler) TickerHand() httpHandler.TicketHandler {
//	return h.ticketHand
//}
//
//func (h *handler) UserHand() httpHandler.UserHandler {
//	return h.userHand
//}

func (h *handler) RegisterAll(ctx context.Context, r *gin.RouterGroup) error {
	api := r.Group("/api")
	{
		h.userHand.RegisterRouter(api)
		h.ticketHand.RegisterRouter(api)
		h.flightHand.RegisterRouter(api)
	}

	return nil
}

func NewHandler(usecase Usecase) (Handler, error) {
	return &handler{
		flightHand: httpHandler.NewFlightHandler(usecase.FlightUc()),
		ticketHand: httpHandler.NewTicketHandler(usecase.TicketUc()),
		userHand:   httpHandler.NewUserHandler(usecase.UserUc()),
	}, nil
}
