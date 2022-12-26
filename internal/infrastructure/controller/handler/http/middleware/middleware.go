package httpMiddleware

import (
	"github.com/achillescres/saina-api/internal/config"
	"github.com/achillescres/saina-api/internal/domain/service"
	"github.com/achillescres/saina-api/pkg/object/oid"
	"github.com/gin-gonic/gin"
)

type Middleware interface {
	ParseAndInjectTokenMiddleware(c *gin.Context)
	GetUserId(c *gin.Context) (oid.Id, error)
}

type middleware struct {
	middlewareConfig config.MiddlewareConfig
	authService      service.AuthService
}

func NewMiddleware(middlewareConfig config.MiddlewareConfig, authService service.AuthService) Middleware {
	return &middleware{middlewareConfig: middlewareConfig, authService: authService}
}
