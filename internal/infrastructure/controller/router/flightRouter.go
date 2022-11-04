package router

import (
	"api-app/internal/infrastructure/controller/handler/http"
	"api-app/pkg/rfmt"
	"github.com/gin-gonic/gin"
)

func flightRouter(handler httpHandler.FlightHandler) *gin.Engine {
	r := gin.Default()

	r.GET(
		rfmt.JoinApi("getAllFlightTables"),
		handler.GetAllFlightTables,
	)

	return r
}
