package service

import (
	"api-app/internal/config"
	"api-app/internal/domain/entity"
	"api-app/internal/domain/object"
	"api-app/internal/domain/storage"
	"api-app/internal/infrastructure/controller/sto"
	"api-app/pkg/object/oid"
	"context"
)

type DataService interface {
	GetAllFlightTables(ctx context.Context) ([]*sto.FlightTableSTO, error)
	GetAllFlightsInMap(ctx context.Context) (map[oid.Id]*entity.Flight, error)
}

type dataService struct {
	flightStorage storage.FlightStorage
	ticketStorage storage.TicketStorage
	cfg           config.TablesConfig
}

var _ DataService = (*dataService)(nil)

func NewDataService(flightStorage storage.FlightStorage, ticketStorage storage.TicketStorage) DataService {
	return &dataService{flightStorage: flightStorage, ticketStorage: ticketStorage}
}

func (dataS *dataService) GetAllFlightTables(ctx context.Context) ([]*sto.FlightTableSTO, error) {
	flights, err := dataS.flightStorage.GetAllInMap(ctx)
	if err != nil {
		return nil, err
	}
	tickets, err := dataS.ticketStorage.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	fTableSTOsMap := map[oid.Id]*sto.FlightTableSTO{}

	for _, ticket := range tickets {
		_, contains := fTableSTOsMap[ticket.FlightId]
		if !contains {
			fTableSTOsMap[ticket.FlightId] = sto.ToFlightTableSTO(object.NewFlightTable(
				*flights[ticket.FlightId],
				dataS.cfg.FlightTableDefaultCapacity,
			))
		}
		fT, _ := fTableSTOsMap[ticket.FlightId]
		fT.Tickets = append(fT.Tickets, *ticket)
	}

	fTs := make([]*sto.FlightTableSTO, 0, dataS.cfg.FlightTableDefaultCapacity)
	for _, fT := range fTableSTOsMap {
		fTs = append(fTs, fT)
	}

	return fTs, nil
}

func (dataS *dataService) GetAllFlightsInMap(ctx context.Context) (map[oid.Id]*entity.Flight, error) {
	return dataS.flightStorage.GetAllInMap(ctx)
}
