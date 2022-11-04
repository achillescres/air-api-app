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
	cfg config.UsecaseConfigInvoker
}

var _ FlightUsecase = (*flightUsecase)(nil)

func NewFlightUsecase(
	flightService service.FlightService,
	ticketService service.TicketService,
	cfg config.UsecaseConfigInvoker,
) *flightUsecase {
	return &flightUsecase{FlightService: flightService, TicketService: ticketService, cfg: cfg}
}

func (fUc *flightUsecase) GetAllFlightTables() map[string]entity.FlightTable {
	flights := fUc.GetAllFlightsMap()
	tickets := fUc.GetAllTickets()

	fTableSTOsMap := map[string]entity.FlightTable{}

	for _, ticket := range tickets {
		_, contains := fTableSTOsMap[ticket.FlightId]
		if !contains {
			fTableSTOsMap[ticket.FlightId] = *entity.NewFlightTable(
				flights[ticket.FlightId],
				fUc.cfg().TableCapacity,
			)
		}
		fT, _ := fTableSTOsMap[ticket.FlightId]
		fT.Tickets = append(fT.Tickets, ticket)
	}

	return fTableSTOsMap
}
