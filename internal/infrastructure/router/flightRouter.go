package router

import (
	"api-app/internal/infrastructure/handler"
	"api-app/pkg/rfmt"
	"github.com/gin-gonic/gin"
)

func flightRouter(handler handler.FlightHandler) *gin.Engine {
	r := gin.Default()

	r.GET(
		rfmt.JoinApi("getAllFlightTables"),
		handler.GetAllFlightTables,
	)

	return r
}
