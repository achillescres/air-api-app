package router

import (
	"api-app/internal/infrastructure/controller/handler/http"
	"api-app/pkg/rfmt"
	"github.com/gin-gonic/gin"
)

func RegisterFlightRouter(r *gin.Engine, handler httpHandler.FlightHandler) {
	r.GET(
		rfmt.JoinApiRoute("getAllFlightTables"),
		handler.GetAllFlightTables,
	)
}
