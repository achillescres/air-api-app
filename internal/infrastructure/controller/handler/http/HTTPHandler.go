package httpHandler

import (
	"api-app/internal/domain/entity"
	"github.com/gin-gonic/gin"
)

type Handler[Entity entity.Entity] interface {
	RegisterRouter(r *gin.RouterGroup)
}
