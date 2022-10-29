package parser

import (
	"api-app/internal/usecase"
	"api-app/internal/usecase/dto/flightDTO"
	"api-app/internal/usecase/dto/ticketDTO"
	"bufio"
	"errors"
	log "github.com/sirupsen/logrus"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type TaisParser interface {
	ParseFile(path string) error
	parseFlightRow(row []string) *flightDTO.CreateFlightDTO
	parseTicketRow(flightId string, row []string) *ticketDTO.CreateTicketDTO
}

type taisParser struct {
	tUsecase usecase.TicketUsecase
	fUsecase usecase.FlightUsecase
}

var _ TaisParser = (*taisParser)(nil)

func NewTaisParser(tUsecase usecase.TicketUsecase, fUsecase usecase.FlightUsecase) *taisParser {
	return &taisParser{tUsecase: tUsecase, fUsecase: fUsecase}
}

func (p *taisParser) parseFlightRow(row []string) *flightDTO.CreateFlightDTO {
	correct := true
	airlCode := row[0]
	fltNum := row[1]
	fltDate := row[2][:7+1]
	origIATA := row[3]
	destIATA := row[4]
	totalCash, err := strconv.ParseFloat(row[9], 32)
	if err != nil {
		correct = false
	}

	return flightDTO.NewCreateFlightDTO(
		airlCode,
		fltNum,
		fltDate,
		origIATA,
		destIATA,
		totalCash,
		correct,
	)
}

func (p *taisParser) parseTicketRow(flightId string, row []string) *ticketDTO.CreateTicketDTO {
	correct := true
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
		log.Errorf("error cant parse ticket amount: %s", err.Error())
		amount = -1
		correct = false
	}
	delimeter := 23
	totalCash, err := strconv.ParseFloat(row[5][delimeter+1:], 32)
	if err != nil {
		log.Errorf("error cant parse ticket total cash: %s", err.Error())
		totalCash = -1
		correct = false
	}

	return ticketDTO.NewCreateTicketDTO(
		flightId,
		airlCode,
		fltNum,
		fltDate,
		ticketCode,
		ticketType,
		amount,
		totalCash,
		correct,
	)
}

func (p *taisParser) ParseFile(path string) error {
	f, err := os.OpenFile(path, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Errorf("error opening file for parse: %s", err.Error())
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
		log.Errorf("error reading file for parse: %s", err.Error())
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
			parsedFlight := p.parseFlightRow(procLine)
			flightId, err = p.fUsecase.CreateFlight(*parsedFlight)
			if err != nil {
				globalErr = err
				log.Fatalf("fatal creating flight: %s", err)
			}
		case 6: // ticket
			parsedTicket := p.parseTicketRow(flightId, procLine)
			err := p.tUsecase.CreateTicket(*parsedTicket)
			if err != nil {
				globalErr = err
				log.Fatalf("fatal creating ticket: %s", err.Error())
			}
		}
	}

	return globalErr
}
