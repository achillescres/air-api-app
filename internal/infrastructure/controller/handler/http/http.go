package httpHandler

import (
	"api-app/internal/config"
	"api-app/internal/domain/service"
	"github.com/gin-gonic/gin"
)

type Handler interface {
	RegisterRouter(r *gin.RouterGroup) error
}

type handler struct {
	flightService service.FlightService
	ticketService service.TicketService
	userService   service.UserService
	cfg           config.HandlerConfig
}

var _ Handler = (*handler)(nil)

func NewHandler(
	flightService service.FlightService,
	ticketService service.TicketService,
	userService service.UserService,
	cfg *config.HandlerConfig,
) Handler {
	return &handler{flightService: flightService, ticketService: ticketService, userService: userService, cfg: *cfg}
}

func (h *handler) RegisterRouter(r *gin.RouterGroup) error {
	auth := r.Group("/auth")
	h.registerUser(auth)

	api := r.Group("/api")
	h.registerFlightTable(api)
	h.registerTicket(api)

	return nil
}
