package usecase

import (
	"api-app/internal/adapter/service"
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
	FlightService service.FlightService
	TicketService service.TicketService
}

var _ FlightUsecase = (*flightUsecase)(nil)

func (fUsecase *flightUsecase) GetFlightById(id string) *flightDTO.ReadFlightDTO {
	//TODO implement me
	panic("implement me")
}

func (fUsecase *flightUsecase) GetAllFlightTables() []*flightTableDTO.ResponseFlightTableDTO {
	flights := fUsecase.FlightService.GetAllFlightsMap()
	tickets := fUsecase.TicketService.GetAllTickets()

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
	err := fUsecase.FlightService.CreateFlight(*flight)
	if err != nil {
		return "", err
	}

	return id, nil
}

func NewFlightUsecase(FlightService service.FlightService, ticketService service.TicketService) *flightUsecase {
	return &flightUsecase{FlightService: FlightService, TicketService: ticketService}
}
