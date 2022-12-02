package service

import (
	"api-app/internal/config"
	"api-app/internal/domain/storage"
	"api-app/internal/domain/storage/dto"
	"api-app/pkg/object/oid"
	"context"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"strconv"
	"unicode"
)

type ParserService interface {
	ParseFields(ctx context.Context, fields []string) error
	//ParseRawRows(ctx context.Context, rrows []string) error
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

// A4 101 2022021312 KRR VKO 19502200 00SU9 0 NN 000151280.00
func (*parserService) parseFlightRow(fields []string) (*dto.FLightCreate, error) {
	if len(fields) != 10 {
		return nil, errors.New("flight fields len must be 10")
	}

	correctlyParsed := true

	airlCode := fields[0]

	fltNum := fields[1]
	fltDate := fields[2][:7+1]

	origIATA := fields[3]
	destIATA := fields[4]

	departureTimeStr := fields[5][:len(fields[5])/2]
	depTHour, err := strconv.Atoi(departureTimeStr[:2])
	depTMinute, err := strconv.Atoi(departureTimeStr[2:])
	departureTime := fmt.Sprintf("%d:%d", depTHour, depTMinute)

	arriveTimeStr := fields[5][len(fields[5])/2:]
	arrTHour, err := strconv.Atoi(arriveTimeStr[:2])
	arrTMinute, err := strconv.Atoi(arriveTimeStr[2:])
	arriveTime := fmt.Sprintf("%d:%d", arrTHour, arrTMinute)

	totalCash, err := strconv.ParseFloat(fields[9], 32)
	if err != nil {
		correctlyParsed = false
	}

	return &dto.FLightCreate{
		AirlCode:        airlCode,
		FltNum:          fltNum,
		FltDate:         fltDate,
		OrigIATA:        origIATA,
		DestIATA:        destIATA,
		DepartureTime:   departureTime,
		ArrivalTime:     arriveTime,
		TotalCash:       totalCash,
		CorrectlyParsed: correctlyParsed,
	}, nil
}

// A4 101 2022021312B Y0100Y 00 0000000001000020000000050000000000000.00
func (pS *parserService) parseTicketRow(flightId oid.Id, fields []string) (*dto.TicketCreate, error) {
	if len(fields) != 6 {
		return nil, errors.New("ticket fields len must be 6")
	}

	correctlyParsed := true
	airlCode := fields[0]
	fltNum := fields[1]
	fltDate := fields[2][:7+1]
	ticketCode := fields[3]
	ticketCapacity, err := strconv.Atoi(fields[3][1:5])
	if err != nil {
		log.Errorf("error cant atoi ticket capacity: %s\n", err.Error())
		ticketCapacity = -1
		correctlyParsed = false
	}
	ticketType := string(ticketCode[len(ticketCode)-1])
	if unicode.IsNumber(rune(ticketType[0])) {
		ticketCode = "official"
	}

	amount, err := strconv.Atoi(fields[5][:3+1])
	if err != nil {
		log.Errorf("error cant parse ticket amount: %s\n", err.Error())
		amount = -1
		correctlyParsed = false
	}

	totalCash, err := strconv.ParseFloat(fields[5][pS.cfg.TotalCashDelimiterIndex+1:], 32)
	if err != nil {
		log.Errorf("error cant parse ticket total cash: %s\n", err.Error())
		totalCash = -1
		correctlyParsed = false
	}

	return &dto.TicketCreate{
		FlightId:        flightId,
		AirlCode:        airlCode,
		FltNum:          fltNum,
		FltDate:         fltDate,
		TicketCode:      ticketCode,
		TicketCapacity:  ticketCapacity,
		TicketType:      ticketType,
		Amount:          amount,
		TotalCash:       totalCash,
		CorrectlyParsed: correctlyParsed,
	}, nil
}

func (pS *parserService) ParseFields(ctx context.Context, fields []string) error {
	switch len(fields) {
	case 10:
		parsedFlight, err := pS.parseFlightRow(fields)
		if err != nil {
			return err
		}

		flight, err := pS.flightStorage.Store(ctx, *parsedFlight)
		pS.currentFlightId = flight.Id
		if err != nil {
			return err
		}
	case 6:
		parsedTicket, err := pS.parseTicketRow(pS.currentFlightId, fields)
		if err != nil {
			return err
		}

		_, err = pS.ticketStorage.Store(ctx, *parsedTicket)
		if err != nil {
			return err
		}
	default:
		return errors.New("flight or ticket fields len must be 10 or 7 respectively")
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
