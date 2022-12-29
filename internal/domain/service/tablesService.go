package service

import (
	"context"
	"github.com/achillescres/saina-api/internal/config"
	"github.com/achillescres/saina-api/internal/domain/entity"
	"github.com/achillescres/saina-api/internal/domain/object"
	storage2 "github.com/achillescres/saina-api/internal/domain/storage"
	"github.com/achillescres/saina-api/internal/infrastructure/controller/safeObject"
	"github.com/achillescres/saina-api/pkg/object/oid"
)

type DataService interface {
	GetAllFlightTables(ctx context.Context) ([]*safeObject.FlightTableSafe, error)
	GetAllFlightsInMap(ctx context.Context) (map[oid.Id]*entity.Flight, error)
}

type dataService struct {
	flightStorage storage2.FlightStorage
	ticketStorage storage2.TicketStorage
	cfg           config.TablesConfig
}

var _ DataService = (*dataService)(nil)

func NewDataService(flightStorage storage2.FlightStorage, ticketStorage storage2.TicketStorage) DataService {
	return &dataService{flightStorage: flightStorage, ticketStorage: ticketStorage}
}

func (dataS *dataService) GetAllFlightTables(ctx context.Context) ([]*safeObject.FlightTableSafe, error) {
	flights, err := dataS.flightStorage.GetAllInMap(ctx)
	if err != nil {
		return nil, err
	}
	tickets, err := dataS.ticketStorage.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	fTableSTOsMap := map[oid.Id]*safeObject.FlightTableSafe{}

	for _, ticket := range tickets {
		_, contains := fTableSTOsMap[ticket.FlightId]
		if !contains {
			fTableSTOsMap[ticket.FlightId] = safeObject.ToFlightTableSafe(object.NewFlightTable(
				*flights[ticket.FlightId],
				dataS.cfg.FlightTableDefaultCapacity,
			))
		}
		fT, _ := fTableSTOsMap[ticket.FlightId]
		fT.Tickets = append(fT.Tickets, *ticket)
	}

	fTs := make([]*safeObject.FlightTableSafe, len(fTableSTOsMap))
	i := 0
	for _, fT := range fTableSTOsMap {
		fTs[i] = fT
		i++
	}

	return fTs, nil
}

func (dataS *dataService) GetAllFlightsInMap(ctx context.Context) (map[oid.Id]*entity.Flight, error) {
	return dataS.flightStorage.GetAllInMap(ctx)
}
