package product

import (
	"context"
	"github.com/gin-gonic/gin"
)

type Routers struct {
	*gin.Engine
}

func NewRouters(ctx context.Context, handlers Handlers) (*Routers, error) {
	r := gin.Default()

	root := r.Group("/")
	err := handlers.RegisterAll(ctx, root)
	if err != nil {
		return nil, err
	}

	return &Routers{Engine: r}, nil
}
