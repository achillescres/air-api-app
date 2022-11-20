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
	Usecase
	GetAllFlightTables(ctx context.Context) (map[oid.Id]*object.FlightTable, error)
	StoreFlight(ctx context.Context, fC dto.FLightCreate) (*entity.Flight, error)
	StoreTicket(ctx context.Context, tC dto.TicketCreate) (*entity.Ticket, error)
}

type flightUsecase struct {
	flightService service.FlightService
	ticketService service.TicketService
	cfg           config.UsecaseConfig
}

var _ FlightUsecase = (*flightUsecase)(nil)

func NewFlightUsecase(
	flightService service.FlightService,
	ticketService service.TicketService,
	cfg config.UsecaseConfig,
) FlightUsecase {
	return &flightUsecase{flightService: flightService, ticketService: ticketService, cfg: cfg}
}

func (fUc *flightUsecase) StoreFlight(ctx context.Context, fC dto.FLightCreate) (*entity.Flight, error) {
	flight, err := fUc.flightService.Store(ctx, fC)
	if err != nil {
		return nil, err
	}
	return flight, nil
}

func (fUc *flightUsecase) StoreTicket(ctx context.Context, tC dto.TicketCreate) (*entity.Ticket, error) {
	store, err := fUc.ticketService.Store(ctx, tC)
	if err != nil {
		return nil, err
	}
	return store, nil
}

func (fUc *flightUsecase) GetAllFlightTables(ctx context.Context) (map[oid.Id]*object.FlightTable, error) {
	flights, err := fUc.flightService.GetAllByMap(ctx)
	if err != nil {
		return nil, err
	}
	tickets, err := fUc.ticketService.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	fTableSTOsMap := map[oid.Id]*object.FlightTable{}

	for _, ticket := range tickets {
		_, contains := fTableSTOsMap[ticket.FlightId]
		if !contains {
			fTableSTOsMap[ticket.FlightId] = object.NewFlightTable(
				*flights[ticket.FlightId],
				fUc.cfg.DefaultTableCapacity,
			)
		}
		fT, _ := fTableSTOsMap[ticket.FlightId]
		fT.Tickets = append(fT.Tickets, *ticket)
	}

	return fTableSTOsMap, nil
}
