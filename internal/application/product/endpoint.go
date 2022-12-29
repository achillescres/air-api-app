package product

import (
	"context"
	"github.com/gin-gonic/gin"
)

// Endpoint is artificial product layer that helps with providing logic to outer scopes
type Endpoint interface {
	RegisterHandlersToGroup(engine *gin.Engine) error
	RunTaisParserController(ctx context.Context) error
}

type endpoint struct {
	controllers Controllers
}

func (e *endpoint) RegisterHandlersToGroup(r *gin.Engine) error {
	root := r.Group("/")
	err := e.controllers.RegisterHandlers(root)
	return err
}

func (e *endpoint) RunTaisParserController(ctx context.Context) error {
	err := e.controllers.RunTaisController(ctx)
	return err
}

func NewEndpoint(cs Controllers) Endpoint {
	return &endpoint{controllers: cs}
}
