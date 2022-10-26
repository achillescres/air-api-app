package usecase

import (
	"api-app/internal/adapter/service"
	"api-app/internal/usecase/composite"
	"api-app/internal/usecase/dto/flightDTO"
	"api-app/internal/usecase/dto/flightTableDTO"
	"github.com/google/uuid"
)

type FlightUsecase interface {
	GetAllFlightTables() []*flightTableDTO.ResponseFlightTableDTO
	GetFlightTableById(id string) *flightTableDTO.ResponseFlightTableDTO
	GetFlightById(id string) *flightDTO.ReadFlightDTO

	CreateFlight(createFlight flightDTO.CreateFlightDTO) (string, error)
}

type flightUsecase struct {
	flightService service.FlightService
	ticketService service.TicketService
}

var _ FlightUsecase = (*flightUsecase)(nil)

func (fUsecase *flightUsecase) GetFlightById(id string) *flightDTO.ReadFlightDTO {
	//TODO implement me
	panic("implement me")
}

func (fUsecase *flightUsecase) GetAllFlightTables() []*flightTableDTO.ResponseFlightTableDTO {
	flights := fUsecase.flightService.GetAllFlightsMap()
	tickets := fUsecase.ticketService.GetAllTickets()

	flightTables := map[string]*flightTableDTO.ResponseFlightTableDTO{}

	for _, ticket := range tickets {
		fT, contains := flightTables[ticket.FlightId]
		if contains {
			fT.Tickets = append(fT.Tickets, ticket)
		} else {
			flightTables[ticket.FlightId] = flightTableDTO.NewResponseFlightTableDTOFromFlight(
				*flights[ticket.FlightId],
				nil,
			)
		}
	}

	fT := make([]*flightTableDTO.ResponseFlightTableDTO, 0)
	for _, v := range flightTables {
		fT = append(fT, v)
	}

	return fT
}

func (fUsecase *flightUsecase) GetFlightTableById(id string) *flightTableDTO.ResponseFlightTableDTO {
	//TODO implement me
	panic("implement me")
}

func (fUsecase *flightUsecase) GetTicketById(id string) *flightDTO.ReadFlightDTO {
	//TODO implement me
	panic("implement me")
}

func (fUsecase *flightUsecase) CreateFlight(createFlight flightDTO.CreateFlightDTO) (string, error) {
	id := uuid.New().String()
	flight := flightDTO.NewFlightFromCreateFlightDTO(id, createFlight)
	err := fUsecase.flightService.CreateFlight(*flight)
	if err != nil {
		return "", err
	}

	return id, nil
}

func NewFlightUsecase(flightService service.FlightService, ticketService service.TicketService) *flightUsecase {
	return &flightUsecase{flightService: flightService, ticketService: ticketService}
}

func GenerateFlightUsecase() *flightUsecase {
	flightComposite := composite.GenerateFlightComposite()
	ticketComposite := composite.GenerateTicketComposite()
	return NewFlightUsecase(flightComposite, ticketComposite)
}
