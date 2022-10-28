package router

import (
	"api-app/internal/infrastructure/handler"
	"github.com/gin-gonic/gin"
)

func ticketRouter(handler handler.TicketHandler) *gin.Engine {
	r := gin.Default()

	return r
}
