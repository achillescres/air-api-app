package product

import (
	"context"
	"github.com/gin-gonic/gin"
)

type Router struct {
	*gin.Engine
}

func NewRouter(ctx context.Context, handlers Handler) (*Router, error) {
	r := gin.Default()

	root := r.Group("/")
	err := handlers.RegisterAll(ctx, root)
	if err != nil {
		return nil, err
	}

	return &Router{Engine: r}, nil
}
