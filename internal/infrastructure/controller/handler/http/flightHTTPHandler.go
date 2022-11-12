package httpHandler

import (
	"api-app/internal/domain/entity"
	"api-app/internal/domain/usecase"
	"api-app/internal/infrastructure/controller/sto"
	"api-app/pkg/object/oid"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type FlightHandler interface {
	Handler[entity.Flight]
	GetAllFlightTables(c *gin.Context)
}

type flightHandler struct {
	flightUsecase usecase.FlightUsecase
}

var _ FlightHandler = (*flightHandler)(nil)

func (fHandler *flightHandler) GetAllFlightTables(c *gin.Context) {
	flightTables, err := fHandler.flightUsecase.GetAllFlightTables()
	if err != nil {
		err := c.AbortWithError(http.StatusInternalServerError, err)
		if err != nil {
			log.Errorf("can't AbortWithError: %s", err.Error())
		}
	}

	flightTableSTOs := map[oid.Id]sto.FlightTableSTO{}
	for _, fT := range flightTables {
		flightTableSTOs[fT.Id] = *sto.ToFLightTableSTO(fT)
	}

	c.JSON(http.StatusOK, flightTableSTOs)
}

func NewFlightHandler(fUc usecase.FlightUsecase) FlightHandler {
	return &flightHandler{flightUsecase: fUc}
}
