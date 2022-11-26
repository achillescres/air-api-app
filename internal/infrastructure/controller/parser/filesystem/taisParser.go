package parser

import (
	"api-app/internal/config"
	"api-app/internal/domain/dto"
	"api-app/internal/domain/service"
	"api-app/pkg/object/oid"
	"bufio"
	"context"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"path"
	"strconv"
	"strings"
	"unicode"
)

type TaisParser interface {
	ParseFirstTaisFile(ctx context.Context) error
}

type taisParser struct {
	tService service.TicketService
	fService service.FlightService
	cfg      config.TaisParserConfig
}

var _ TaisParser = (*taisParser)(nil)

func NewTaisParser(
	tService service.TicketService,
	fService service.FlightService,
	cfg config.TaisParserConfig,
) TaisParser {
	return &taisParser{tService: tService, fService: fService, cfg: cfg}
}

// A4 101 2022021312 KRR VKO 19502200 00SU9 0 NN 000151280.00
func parseFlightRow(row []string) *dto.FLightCreate {
	correctlyParsed := true

	airlCode := row[0]

	fltNum := row[1]
	fltDate := row[2][:7+1]

	origIATA := row[3]
	destIATA := row[4]

	departureTimeStr := row[5][:len(row[5])/2]
	depTHour, err := strconv.Atoi(departureTimeStr[:2])
	depTMinute, err := strconv.Atoi(departureTimeStr[2:])
	departureTime := fmt.Sprintf("%d:%d", depTHour, depTMinute)

	arriveTimeStr := row[5][len(row[5])/2:]
	arrTHour, err := strconv.Atoi(arriveTimeStr[:2])
	arrTMinute, err := strconv.Atoi(arriveTimeStr[2:])
	arriveTime := fmt.Sprintf("%d:%d", arrTHour, arrTMinute)

	totalCash, err := strconv.ParseFloat(row[9], 32)
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
func (taisPrsr *taisParser) parseTicketRow(flightId oid.Id, row []string) *dto.TicketCreate {
	correctlyParsed := true
	airlCode := row[0]
	fltNum := row[1]
	fltDate := row[2][:7+1]
	ticketCode := row[3]
	ticketCapacity, err := strconv.Atoi(row[3][1:5])
	if err != nil {
		log.Errorf("error cant atoi ticket capacity: %s\n", err.Error())
		ticketCapacity = -1
		correctlyParsed = false
	}
	ticketType := string(ticketCode[len(ticketCode)-1])
	if unicode.IsNumber(rune(ticketType[0])) {
		ticketCode = "official"
	}

	amount, err := strconv.Atoi(row[5][:3+1])
	if err != nil {
		log.Errorf("error cant parse ticket amount: %s\n", err.Error())
		amount = -1
		correctlyParsed = false
	}

	totalCash, err := strconv.ParseFloat(row[5][taisPrsr.cfg.TotalCashDelimiterIndex+1:], 32)
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
	}
}

func (taisPrsr *taisParser) ParseFirstTaisFile(ctx context.Context) error {
	env := config.Env()
	taisDirPath := path.Join(env.ProjectAbsPath, taisPrsr.cfg.TaisDirPath)
	inDir, err := os.ReadDir(taisDirPath)
	if err != nil {
		// TODO remove fatal drop, add logic to save the system from deprecated data(outer layer work)
		log.Fatalf("(WillRemLog)fatal scanning tais directory=%s: %s\n", taisDirPath, err.Error())
		return err
	}

	taisFileName := ""
	for _, entry := range inDir {
		if !entry.IsDir() {
			taisFileName = entry.Name()
		}
	}

	if taisFileName == "" {
		log.Fatalf("(WillRemLog)fatal didnt find tais file in tais dir=%s\n", taisDirPath) // TODO remove fatals
		return errors.New(fmt.Sprintf("error didnt find tais file in tais dir=%s\n", taisDirPath))
	}

	taisFilePath := path.Join(taisDirPath, taisFileName)
	f, err := os.OpenFile(taisFilePath, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("(WillRemLog)error opening file for parse: %s\n", err.Error()) // TODO remove fatals
		return err
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			// TODO add logic to prevent memory leaks from unclosed files
			log.Errorf("error closing tais file=%s: %s\n", f.Name(), err.Error())
		}
	}(f)

	sc := bufio.NewScanner(f)
	sc.Split(bufio.ScanLines)
	if !sc.Scan() {
		log.Errorln("error tais parse file is empty")
		return errors.New("error parse file is empty")
	} else {
		sc.Text() // Meta line
	}

	rows := make([]string, 0, 1800)
	for sc.Scan() {
		rows = append(rows, sc.Text())
	}
	if err := sc.Err(); err != nil {
		log.Fatalf("(WillRemLog)error reading tais file for parse: %s\n", err.Error()) // TODO remove fatals
		return err
	}

	var (
		flightId  oid.Id
		globalErr error = nil
	)
	for _, row := range rows {
		procLine := strings.Fields(strings.TrimSpace(row))
		switch len(procLine) {
		case 10: // flight
			parsedFlight := parseFlightRow(procLine)
			flight, err := taisPrsr.fService.Store(ctx, *parsedFlight)
			flightId = flight.Id
			if err != nil {
				globalErr = err
				log.Fatalf("(WillRemLog)fatal creating flight: %s\n", err) // TODO remove fatals
				return err
			}
		case 6: // ticket
			parsedTicket := taisPrsr.parseTicketRow(flightId, procLine)
			_, err := taisPrsr.tService.Store(ctx, *parsedTicket)
			if err != nil {
				globalErr = err
				log.Fatalf("(WillRemLog)fatal creating ticket: %s\n", err.Error()) // TODO remove fatals
				return err
			}
		}
	}

	return globalErr
}
