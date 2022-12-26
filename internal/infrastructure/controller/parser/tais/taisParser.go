package parser

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"github.com/achillescres/saina-api/internal/config"
	"github.com/achillescres/saina-api/internal/domain/service"
	dto "github.com/achillescres/saina-api/internal/domain/storage/dto"
	"github.com/achillescres/saina-api/pkg/object/oid"
	log "github.com/sirupsen/logrus"
	"os"
	"path"
	"strconv"
	"strings"
	"unicode"
)

type TaisParser interface {
	ParseFirstTaisFile(ctx context.Context) (*string, *TaisFileParsingErrors, error)
}

type taisParser struct {
	parserService service.ParserService
	cfg           config.TaisParserConfig
}

type TaisFileParsingErrors struct {
	Meta string
	Errs map[int]error
}

func NewTaisParser(
	parserService service.ParserService,
	cfg config.TaisParserConfig,
) TaisParser {
	return &taisParser{parserService: parserService, cfg: cfg}
}

var _ TaisParser = (*taisParser)(nil)

func (tP *taisParser) ParseFirstTaisFile(ctx context.Context) (*string, *TaisFileParsingErrors, error) {
	env := config.Env()
	taisDirPath := path.Join(env.ProjectAbsPath, tP.cfg.TaisDirPath)
	dirEntries, err := os.ReadDir(taisDirPath)
	if err != nil {
		log.Errorf("error (TaisParser) scanning tais directory=%s: %s\n", taisDirPath, err)
		return nil, nil, err
	}

	taisFileName := ""
	for _, entry := range dirEntries {
		if !entry.IsDir() && strings.HasPrefix(strings.ToLower(entry.Name()), "tais") {
			taisFileName = entry.Name()
		}
	}

	if taisFileName == "" {
		log.Errorf("error (TaisParser) didnt find tais file in tais dir=%s\n", taisDirPath)
		return nil, nil, errors.New(
			fmt.Sprintf("error (TaisParser) didnt find tais file in tais dir=%s", taisDirPath),
		)
	}

	taisFilePath := path.Join(taisDirPath, taisFileName)
	f, err := os.OpenFile(taisFilePath, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Errorf("error (TaisParser) opening file for parse: %s\n", err)
		return nil, nil, err
	}

	defer func() {
		err := f.Close()
		if err != nil {
			log.Errorf("error closing tais file=%s: %s\n", f.Name(), err)
		}
	}()

	sc := bufio.NewScanner(f)
	sc.Split(bufio.ScanLines)
	if !sc.Scan() {
		log.Errorln("error tais parse file is empty")
		return nil, nil, errors.New("error parse file is empty")
	}

	meta := strings.TrimSpace(sc.Text()) // Tais file meta header

	rows := make([]string, 0, tP.cfg.DefaultFlightsSliceCapacity)
	for sc.Scan() {
		rows = append(rows, sc.Text())
	}
	if err := sc.Err(); err != nil {
		log.Fatalf("error (TaisParser) reading tais file for parse: %s\n", err) // TODO remove fatals
		return nil, nil, err
	}

	currentFlightIsCorrupted := false
	flightId := oid.Undefined
	flightNum := ""
	errs := new(TaisFileParsingErrors)
	errs.Meta = meta
	for i, row := range rows {
		fields := strings.Fields(strings.TrimSpace(row))
		switch len(fields) {
		case 10: // Flight
			currentFlightIsCorrupted = false
			fC := tP.parseFlightRow(*(*[10]string)(fields))
			flightId, flightNum, err = tP.parserService.PassFlight(ctx, fC)
			if err != nil {
				log.Errorf("error storing new flight: %s", err)
				errs.Errs[i] = err
			}
		case 6: // Ticket
			if !currentFlightIsCorrupted {
				continue
			}
			tC := tP.parseTicketRow(flightId, flightNum, *(*[6]string)(fields))
			err = tP.parserService.PassTicket(ctx, tC)
			if err != nil {
				log.Errorf("error (TaisParser) passing flight to service: %s\n", err)
				errs.Errs[i] = err
			}
		}
	}

	return &meta, errs, nil // TODO implement errs usage
}

// A4 101 2022021312 KRR VKO 19502200 00SU9 0 NN 000151280.00
func (tP *taisParser) parseFlightRow(fields [10]string) *dto.FLightCreate {
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
	}
}

// A4 101 2022021312B Y0100Y 00 0000000001000020000000050000000000000.00
func (tP *taisParser) parseTicketRow(flightId oid.Id, flightNum string, fields [6]string) *dto.TicketCreate {
	correctlyParsed := true
	airlCode := fields[0]
	fltNum := fields[1]
	fltDate := fields[2][:7+1]
	ticketCode := fields[3]
	ticketCapacity, err := strconv.Atoi(fields[3][1:5])
	if err != nil {
		log.Errorf("error cant atoi ticket capacity: %s\n", err)
		ticketCapacity = -1
		correctlyParsed = false
	}
	ticketType := string(ticketCode[len(ticketCode)-1])
	if unicode.IsNumber(rune(ticketType[0])) {
		ticketCode = "official"
	}

	amount, err := strconv.Atoi(fields[5][:3+1])
	if err != nil {
		log.Errorf("error cant parse ticket amount: %s\n", err)
		amount = -1
		correctlyParsed = false
	}

	totalCash, err := strconv.ParseFloat(fields[5][tP.cfg.TotalCashDelimiterIndex+1:], 32)
	if err != nil {
		log.Errorf("error cant parse ticket total cash: %s\n", err)
		totalCash = -1
		correctlyParsed = false
	}

	return dto.NewTicketCreate(
		flightId,
		flightNum,
		airlCode,
		fltNum,
		fltDate,
		ticketCode,
		ticketCapacity,
		ticketType,
		amount,
		totalCash,
		correctlyParsed,
	)
}
