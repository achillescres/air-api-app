package httpHandler

import (
	"github.com/achillescres/saina-api/pkg/gin/ginresponse"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func (h *handler) GetAllTicketsMap(c *gin.Context) {
	ticketsMap, err := h.dataService.GetAllFlightsInMap(c)
	if err != nil {
		log.Errorf("error TicketUsecase.GetAllTicketsMap: %s\n", err)
		ginresponse.Error(c, http.StatusInternalServerError, err, "couldn't get all TicketsMap")
		return
	}

	c.JSON(http.StatusOK, ticketsMap)
}

func (h *handler) registerTicket(r *gin.RouterGroup) {
	r = r.Group("/ticket")
	r.GET("/getAllTicketsMap", h.GetAllTicketsMap)
}
