package httpHandler

import (
	"api-app/internal/domain/object"
	"api-app/internal/infrastructure/controller/sto"
	"api-app/pkg/gin/ginresponse"
	"api-app/pkg/object/oid"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *handler) getAllFlightTables(ctx context.Context) (map[oid.Id]*sto.FlightTableSTO, error) {
	flights, err := h.flightService.GetAllByMap(ctx)
	if err != nil {
		return nil, err
	}
	tickets, err := h.ticketService.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	fTableSTOsMap := map[oid.Id]*sto.FlightTableSTO{}

	for _, ticket := range tickets {
		_, contains := fTableSTOsMap[ticket.FlightId]
		if !contains {
			fTableSTOsMap[ticket.FlightId] = sto.ToFlightTableSTO(object.NewFlightTable(
				*flights[ticket.FlightId],
				h.cfg.DefaultTableCapacity,
			))
		}
		fT, _ := fTableSTOsMap[ticket.FlightId]
		fT.Tickets = append(fT.Tickets, *ticket)
	}

	return fTableSTOsMap, nil
}

func (h *handler) GetAllFlightTables(c *gin.Context) {
	flightTableSTOs, err := h.getAllFlightTables(c)
	if err != nil {
		ginresponse.WithError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, flightTableSTOs) // TODO maybe make it returns list instead of a map
}

func (h *handler) registerFlightTable(r *gin.RouterGroup) {
	r.GET("/getAllFlightTables", h.GetAllFlightTables)
}
