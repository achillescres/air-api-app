package product

import (
	"github.com/gin-gonic/gin"
)

type Routers struct {
	*gin.Engine
}

func NewRouters(handlers Controllers) (*Routers, error) {
	r := gin.Default()

	root := r.Group("/")
	err := handlers.Register(root)
	if err != nil {
		return nil, err
	}

	return &Routers{Engine: r}, nil
}
