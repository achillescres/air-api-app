package httpHandler

import (
	"api-app/internal/domain/entity"
	"api-app/internal/domain/usecase"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type TicketHandler interface {
	Handler[entity.Ticket]
	GetAllTicketsMap(c *gin.Context)
}

type ticketHandler struct {
	ticketUsecase usecase.TicketUsecase
}

var _ TicketHandler = (*ticketHandler)(nil)

func NewTicketHandler(tUc usecase.TicketUsecase) TicketHandler {
	return &ticketHandler{ticketUsecase: tUc}
}

func (tH *ticketHandler) GetAllTicketsMap(c *gin.Context) {
	ticketsMap, err := tH.ticketUsecase.GetAllByMap(c)
	if err != nil {
		log.Errorf("error TicketUsecase.GetAllTicketsMap: %s", err.Error())
		err = c.Error(err)
		if err != nil {
			log.Errorf("error puting error to gin context: %s", err.Error())
		}
	}

	c.JSON(http.StatusOK, ticketsMap)
}

func (tH *ticketHandler) RegisterRouter(r *gin.RouterGroup) {
	r = r.Group("/ticket")
	{
		r.GET("/getAllTicketsMap", tH.GetAllTicketsMap)
	}
}
