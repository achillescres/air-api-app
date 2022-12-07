package service

import (
	"context"
	"github.com/achillescres/saina-api/internal/config"
	"github.com/achillescres/saina-api/internal/domain/storage"
	"github.com/achillescres/saina-api/internal/domain/storage/dto"
	"github.com/achillescres/saina-api/pkg/object/oid"
)

type ParserService interface {
	PassFlight(ctx context.Context, fC dto.FLightCreate) error
	PassTicket(ctx context.Context, tC dto.TicketCreate) error
}

type parserService struct {
	flightStorage   storage.FlightStorage
	ticketStorage   storage.TicketStorage
	cfg             *config.TaisParserConfig
	currentFlightId oid.Id
}

func NewParserService(
	flightStorage storage.FlightStorage,
	ticketStorage storage.TicketStorage,
	cfg *config.TaisParserConfig,
) ParserService {
	return &parserService{flightStorage: flightStorage, ticketStorage: ticketStorage, cfg: cfg, currentFlightId: oid.Undefined}
}

var _ ParserService = (*parserService)(nil)

func (pS *parserService) PassFlight(ctx context.Context, fC dto.FLightCreate) error {
	flight, err := pS.flightStorage.Store(ctx, fC)
	pS.currentFlightId = flight.Id
	if err != nil {
		return err
	}

	return nil
}

func (pS *parserService) PassTicket(ctx context.Context, tC dto.TicketCreate) error {
	// TODO flightId
	_, err := pS.ticketStorage.Store(ctx, tC)
	if err != nil {
		return err
	}

	return nil
}

//
//func (pS *parserService) ParseRawRows(ctx context.Context, rrows []string) error {
//	var gErr []error
//	for _, rrow := range rrows {
//		row := strings.Fields(strings.TrimSpace(rrow))
//		err := pS.parseFields(ctx, row)
//		if err != nil {
//			gErr =
//		}
//	}
//
//	return gErr
//}
