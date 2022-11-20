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
	flightTables, err := fHandler.flightUsecase.GetAllFlightTables(c)
	if err != nil {
		err := c.AbortWithError(http.StatusInternalServerError, err)
		if err != nil {
			log.Errorf("can't AbortWithError: %s\n", err.Error())
		}
		return
	}

	flightTableSTOs := map[oid.Id]*sto.FlightTableSTO{}
	for _, fT := range flightTables {
		flightTableSTOs[fT.Id] = sto.ToFlightTableSTO(*fT)
	}

	c.JSON(http.StatusOK, flightTableSTOs) // TODO maybe make it returns list instead of a map
}

func (fHandler *flightHandler) RegisterRouter(r *gin.RouterGroup) {
	r = r.Group("/flight")
	{
		r.GET("getAllFlightTables", fHandler.GetAllFlightTables)
	}
}

func NewFlightHandler(fUc usecase.FlightUsecase) FlightHandler {
	return &flightHandler{flightUsecase: fUc}
}
