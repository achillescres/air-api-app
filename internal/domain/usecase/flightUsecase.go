package usecase

import (
	"api-app/internal/config"
	"api-app/internal/domain/entity"
	"api-app/internal/domain/service"
)

type FlightUsecase interface {
	service.FlightService
	service.TicketService
	GetAllFlightTables() map[string]entity.FlightTable
}

var _ Usecase = (*FlightUsecase)(nil)

type flightUsecase struct {
	service.FlightService
	service.TicketService
	cfg config.UsecaseConfig
}

var _ FlightUsecase = (*flightUsecase)(nil)

func NewFlightUsecase(
	flightService service.FlightService,
	ticketService service.TicketService,
	cfg config.UsecaseConfig,
) FlightUsecase {
	return &flightUsecase{FlightService: flightService, TicketService: ticketService, cfg: cfg}
}

func (fUc *flightUsecase) GetAllFlightTables() map[string]entity.FlightTable {
	flights := fUc.GetAllFlightsMap()
	tickets := fUc.GetAllTickets()

	fTableSTOsMap := map[string]entity.FlightTable{}

	for _, ticket := range tickets {
		_, contains := fTableSTOsMap[ticket.View.FlightId]
		if !contains {
			fTableSTOsMap[ticket.View.FlightId] = *entity.NewFlightTable(
				flights[ticket.View.FlightId],
				fUc.cfg.DefaultTableCapacity,
			)
		}
		fT, _ := fTableSTOsMap[ticket.View.FlightId]
		fT.Tickets = append(fT.Tickets, ticket)
	}

	return fTableSTOsMap
}
