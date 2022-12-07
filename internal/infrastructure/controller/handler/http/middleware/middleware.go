package httpMiddleware

import (
	"github.com/achillescres/saina-api/internal/config"
	"github.com/achillescres/saina-api/internal/domain/service"
	"github.com/gin-gonic/gin"
)

type Middleware interface {
	ParseAndInjectTokenMiddleware(c *gin.Context)
}

type middleware struct {
	middlewareConfig config.MiddlewareConfig
	authService      service.AuthService
}

func NewMiddleware(middlewareConfig config.MiddlewareConfig, authService service.AuthService) Middleware {
	return &middleware{middlewareConfig: middlewareConfig, authService: authService}
}
