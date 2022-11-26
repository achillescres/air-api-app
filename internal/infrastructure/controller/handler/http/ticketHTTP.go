package httpHandler

import (
	"api-app/pkg/gin/ginresponse"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func (h *handler) GetAllTicketsMap(c *gin.Context) {
	ticketsMap, err := h.ticketService.GetAllByMap(c)
	if err != nil {
		log.Errorf("error TicketUsecase.GetAllTicketsMap: %s\n", err.Error())
		ginresponse.WithError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, ticketsMap)
}

func (h *handler) registerTicket(r *gin.RouterGroup) {
	r.GET("/getAllTicketsMap", h.GetAllTicketsMap)
}
