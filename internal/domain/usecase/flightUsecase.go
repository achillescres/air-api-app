package usecase

import (
	"api-app/internal/config"
	"api-app/internal/domain/entity"
	"api-app/internal/domain/object"
	"api-app/internal/domain/service"
	"api-app/internal/domain/storage/dto"
	"api-app/pkg/object/oid"
	"context"
)

type FlightUsecase interface {
	Usecase[entity.Flight, entity.FlightView, dto.FLightCreate]
	GetAllFlightTables(ctx context.Context) (map[oid.Id]object.FlightTable, error)
}

type flightUsecase struct {
	service.FlightService
	service.TicketService
	cfg config.UsecaseConfig
}

var _ FlightUsecase = (*flightUsecase)(nil)

func (fUc *flightUsecase) GetAllFlightTables(ctx context.Context) (map[oid.Id]object.FlightTable, error) {
	flights, err := fUc.FlightService.GetAllByMap(ctx)
	if err != nil {
		return nil, err
	}
	tickets, err := fUc.TicketService.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	fTableSTOsMap := map[oid.Id]object.FlightTable{}

	for _, ticket := range tickets {
		_, contains := fTableSTOsMap[ticket.View.FlightId]
		if !contains {
			fTableSTOsMap[ticket.View.FlightId] = *object.NewFlightTable(
				flights[ticket.View.FlightId],
				fUc.cfg.DefaultTableCapacity,
			)
		}
		fT, _ := fTableSTOsMap[ticket.View.FlightId]
		fT.Tickets = append(fT.Tickets, ticket)
	}

	return fTableSTOsMap, nil
}

func NewFlightUsecase(
	flightService service.FlightService,
	ticketService service.TicketService,
	cfg config.UsecaseConfig,
) FlightUsecase {
	return &flightUsecase{FlightService: flightService, TicketService: ticketService, cfg: cfg}
}
