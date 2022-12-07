package product

import (
	"github.com/achillescres/saina-api/internal/config"
	"github.com/achillescres/saina-api/internal/domain/service"
	httpMiddleware "github.com/achillescres/saina-api/internal/infrastructure/controller/handler/http/middleware"
)

type Middlewares struct {
	middleware httpMiddleware.Middleware
}

func NewMiddlewares(cfg config.MiddlewareConfig, authService service.AuthService) *Middlewares {
	return &Middlewares{middleware: httpMiddleware.NewMiddleware(cfg, authService)}
}
