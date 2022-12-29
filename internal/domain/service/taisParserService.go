package service

import (
	"context"
	"github.com/achillescres/saina-api/internal/config"
	"github.com/achillescres/saina-api/internal/domain/storage"
	"github.com/achillescres/saina-api/internal/domain/storage/dto"
	"github.com/achillescres/saina-api/pkg/object/oid"
)

type TaisParserService interface {
	StoreFlight(ctx context.Context, fC *dto.FLightCreate) (oid.Id, string, error)
	StoreTicket(ctx context.Context, tC *dto.TicketCreate) error
}

type taisParserService struct {
	flightStorage storage.FlightStorage
	ticketStorage storage.TicketStorage
	cfg           config.TaisConfig
}

func NewParserService(
	flightStorage storage.FlightStorage,
	ticketStorage storage.TicketStorage,
	cfg config.TaisConfig,
) TaisParserService {
	return &taisParserService{flightStorage: flightStorage, ticketStorage: ticketStorage, cfg: cfg}
}

var _ TaisParserService = (*taisParserService)(nil)

func (pS *taisParserService) StoreFlight(ctx context.Context, fC *dto.FLightCreate) (oid.Id, string, error) {
	flight, err := pS.flightStorage.Store(ctx, *fC)
	if err != nil {
		return "", "", err
	}

	return flight.Id, flight.FltNum, nil
}

func (pS *taisParserService) StoreTicket(ctx context.Context, tC *dto.TicketCreate) error {
	// TODO flightId
	_, err := pS.ticketStorage.Store(ctx, *tC)
	if err != nil {
		return err
	}

	return nil
}

//
//func (pS *taisParserService) ParseRawRows(ctx context.Context, rrows []string) error {
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
