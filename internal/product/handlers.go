package product

import (
	"api-app/internal/config"
	httpHandler "api-app/internal/infrastructure/controller/handler/http"
	"github.com/gin-gonic/gin"
)

type Handlers interface {
	Register(r *gin.RouterGroup) error
}

type handlers struct {
	handler httpHandler.Handler
}

func NewHandlers(services Services, cfg *config.HandlerConfig) (Handlers, error) {
	return &handlers{handler: httpHandler.NewHandler(
		services.FlightService(),
		services.TicketService(),
		services.UserService(),
		cfg,
	)}, nil
}

func (h *handlers) Register(r *gin.RouterGroup) error {
	return h.handler.RegisterRouter(r)
}
