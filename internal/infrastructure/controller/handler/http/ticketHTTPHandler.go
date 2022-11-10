package httpHandler

import (
	"api-app/internal/domain/usecase"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type TicketHandler interface {
	GetAllTicketsMap(c *gin.Context)
}

var _ Handler = (*TicketHandler)(nil)

type ticketHandler struct {
	ticketUsecase usecase.TicketUsecase
}

func (tH *ticketHandler) GetAllTicketsMap(c *gin.Context) {
	ticketsMap, err := tH.ticketUsecase.GetAllTicketsMap()
	if err != nil {
		log.Errorf("error TicketUsecase.GetAllTicketsMap: %s", err.Error())
		err = c.Error(err)
		if err != nil {
			log.Errorf("error puting error to gin context: %s", err.Error())
		}

	}

	c.JSON(http.StatusOK, ticketsMap)
}

var _ TicketHandler = (*ticketHandler)(nil)

func NewTicketHandler(tUc usecase.TicketUsecase) *ticketHandler {
	return &ticketHandler{ticketUsecase: tUc}
}
