package router

import (
	"api-app/internal/infrastructure/controller/handler/http"
	"github.com/gin-gonic/gin"
)

func ticketRouter(handler httpHandler.TicketHandler) *gin.Engine {
	r := gin.Default()

	return r
}
