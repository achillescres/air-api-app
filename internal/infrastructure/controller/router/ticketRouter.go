package router

import (
	"api-app/internal/infrastructure/controller/handler/http"
	"api-app/pkg/rfmt"
	"github.com/gin-gonic/gin"
)

func RegisterTicketRouter(r *gin.Engine, handler httpHandler.TicketHandler) {
	r.GET(
		rfmt.JoinApiRoute("getAllTicketsMap"),
		handler.GetAllTicketsMap,
	)
}
