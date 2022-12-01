package httpHandler

import (
	"api-app/pkg/gin/ginresponse"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *handler) GetAllFlightTables(c *gin.Context) {
	flightTableSTOs, err := h.dataService.GetAllFlightTables(c)
	if err != nil {
		ginresponse.WithError(c, http.StatusInternalServerError, err, "couldn't get all FlightTables")
		return
	}

	c.JSON(http.StatusOK, flightTableSTOs)
}

func (h *handler) registerFlightTable(r *gin.RouterGroup) {
	r.GET("/getAllFlightTables", h.GetAllFlightTables)
}
