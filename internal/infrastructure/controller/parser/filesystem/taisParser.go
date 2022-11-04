package parser

import (
	"api-app/internal/config"
	"api-app/internal/domain/entity"
	"api-app/internal/domain/usecase"
	"bufio"
	"errors"
	log "github.com/sirupsen/logrus"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type TaisParser interface {
	ParseTaisFile() error
}

type taisParser struct {
	tUsecase usecase.TicketUsecase
	fUsecase usecase.FlightUsecase
	cfg      config.ParserConfigInvoke
}

func NewTaisParser(tUsecase usecase.TicketUsecase, fUsecase usecase.FlightUsecase, cfg config.ParserConfigInvoke) *taisParser {
	return &taisParser{tUsecase: tUsecase, fUsecase: fUsecase, cfg: cfg}
}

var _ TaisParser = (*taisParser)(nil)

func parseFlightRow(row []string) entity.FlightView {
	correctlyParsed := true
	airlCode := row[0]
	fltNum := row[1]
	fltDate := row[2][:7+1]
	origIATA := row[3]
	destIATA := row[4]
	totalCash, err := strconv.ParseFloat(row[9], 32)
	if err != nil {
		correctlyParsed = false
	}

	return entity.FlightView{
		AirlCode:        airlCode,
		FltNum:          fltNum,
		FltDate:         fltDate,
		OrigIATA:        origIATA,
		DestIATA:        destIATA,
		TotalCash:       totalCash,
		CorrectlyParsed: correctlyParsed,
	}
}

func parseTicketRow(flightId string, row []string) entity.TicketView {
	correctlyParsed := true
	airlCode := row[0]
	fltNum := row[1]
	fltDate := row[2][:7+1]
	ticketCode := row[3]
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
	delimeter := 23
	totalCash, err := strconv.ParseFloat(row[5][delimeter+1:], 32)
	if err != nil {
		log.Errorf("error cant parse ticket total cash: %s\n", err.Error())
		totalCash = -1
		correctlyParsed = false
	}

	return entity.TicketView{
		FlightId:        flightId,
		AirlCode:        airlCode,
		FltNum:          fltNum,
		FltDate:         fltDate,
		TicketCode:      ticketCode,
		TicketType:      ticketType,
		Amount:          amount,
		TotalCash:       totalCash,
		CorrectlyParsed: correctlyParsed,
	}
}

func (p *taisParser) ParseTaisFile() error {
	path := p.cfg().TaisFilePath

	f, err := os.OpenFile(path, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Errorf("error opening file for parse: %s\n", err.Error())
		return err
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	if !sc.Scan() {
		return errors.New("error parse file is empty")
	} else {
		sc.Text() // Meta line
	}

	rows := make([]string, 0)
	for sc.Scan() {
		rows = append(rows, sc.Text())
	}
	if err := sc.Err(); err != nil {
		log.Errorf("error reading file for parse: %s\n", err.Error())
		return err
	}

	var (
		flightId  string
		globalErr error = nil
	)
	for _, row := range rows {
		procLine := strings.Fields(strings.TrimSpace(row))
		switch len(procLine) {
		case 10: // flight
			parsedFlight := parseFlightRow(procLine)
			flight, err := p.fUsecase.CreateFlight(parsedFlight)
			flightId = flight.Id
			if err != nil {
				globalErr = err
				log.Fatalf("fatal creating flight: %s\n", err)
			}
		case 6: // ticket
			parsedTicket := parseTicketRow(flightId, procLine)
			err := p.tUsecase.CreateTicket(parsedTicket)
			if err != nil {
				globalErr = err
				log.Fatalf("fatal creating ticket: %s\n", err.Error())
			}
		}
	}

	return globalErr
}
