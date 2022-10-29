package usecase

import (
	"api-app/internal/domain/entity"
	"api-app/internal/domain/service"
	"api-app/internal/usecase/dto/flightDTO"
	"api-app/internal/usecase/dto/flightTableDTO"
	"github.com/google/uuid"
)

type FlightUsecase interface {
	GetAllFlightTables() []flightTableDTO.ResponseFlightTableDTO
	GetFlightTableById(id string) flightTableDTO.ResponseFlightTableDTO
	GetFlightById(id string) flightDTO.ReadFlightDTO

	CreateFlight(createFlight flightDTO.CreateFlightDTO) (string, error)
}

type flightUsecase struct {
	FlightService service.FlightService
	TicketService service.TicketService
}

var _ FlightUsecase = (*flightUsecase)(nil)

func (fUsecase *flightUsecase) GetFlightById(id string) flightDTO.ReadFlightDTO {
	//TODO implement me
	panic("implement me")
}

func (fUsecase *flightUsecase) GetAllFlightTables() []flightTableDTO.ResponseFlightTableDTO {
	flights := fUsecase.FlightService.GetAllFlightsMap()
	tickets := fUsecase.TicketService.GetAllTickets()

	fTsMap := map[string]*flightTableDTO.ResponseFlightTableDTO{}

	for _, ticket := range tickets {
		_, contains := fTsMap[ticket.FlightId]
		if !contains {
			fTsMap[ticket.FlightId] = flightTableDTO.NewResponseFlightTableDTOFromFlight(
				flights[ticket.FlightId],
				make([]entity.Ticket, 0),
			)
		}
		fT, _ := fTsMap[ticket.FlightId]
		fT.Tickets = append(fT.Tickets, ticket)
	}

	fTs := make([]flightTableDTO.ResponseFlightTableDTO, 0)
	for _, v := range fTsMap {
		fTs = append(fTs, *v)
	}

	return fTs
}

func (fUsecase *flightUsecase) GetFlightTableById(id string) flightTableDTO.ResponseFlightTableDTO {
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
