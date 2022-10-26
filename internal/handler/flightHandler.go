package handler

import (
	"api-app/internal/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type FlightHandler interface {
	GetAllFlightTables(c *gin.Context)
}

type flightHandler struct {
	uc usecase.FlightUsecase
}

var _ FlightHandler = (*flightHandler)(nil)

func (fHandler *flightHandler) GetAllFlightTables(c *gin.Context) {
	flightTables := fHandler.uc.GetAllFlightTables()
	c.JSON(http.StatusOK, flightTables)
}

func NewFlightHandler(uc usecase.FlightUsecase) *flightHandler {
	return &flightHandler{uc: uc}
}
