package product

import (
	"api-app/internal/config"
	httpHandler "api-app/internal/infrastructure/controller/handler/http"
	parser "api-app/internal/infrastructure/controller/parser/filesystem"
	"github.com/gin-gonic/gin"
)

type Handlers interface {
	Register(r *gin.RouterGroup) error
}

type handlers struct {
	handler httpHandler.Handler
}

func NewHandlers(services *Services, cfg *config.HandlerConfig, taisParser parser.TaisParser) (Handlers, error) {
	return &handlers{handler: httpHandler.NewHandler(
		services.AuthService,
		taisParser,
		services.TablesService,
		*cfg,
	)}, nil
}

func (h *handlers) Register(r *gin.RouterGroup) error {
	return h.handler.RegisterRouter(r)
}
