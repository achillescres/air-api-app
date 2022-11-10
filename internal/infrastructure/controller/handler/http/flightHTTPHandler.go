package httpHandler

import (
	"api-app/internal/domain/usecase"
	"api-app/internal/infrastructure/controller/sto"
	"github.com/gin-gonic/gin"
	"net/http"
)

type FlightHandler interface {
	GetAllFlightTables(c *gin.Context)
}

var _ Handler = (*FlightHandler)(nil)

type flightHandler struct {
	flightUsecase usecase.FlightUsecase
}

func NewFlightHandler(fUc usecase.FlightUsecase) FlightHandler {
	return &flightHandler{flightUsecase: fUc}
}

var _ FlightHandler = (*flightHandler)(nil)

func (fHnd *flightHandler) GetAllFlightTables(c *gin.Context) {
	flightTables := fHnd.flightUsecase.GetAllFlightTables()
	flightTableSTOs := map[string]sto.FlightTableSTO{}
	for _, fT := range flightTables {
		flightTableSTOs[fT.Id] = *sto.ToFLightTableSTO(fT)
	}

	c.JSON(http.StatusOK, flightTableSTOs)
}
