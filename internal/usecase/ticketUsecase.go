package usecase

import (
	"api-app/internal/domain/service"
	"api-app/internal/usecase/dto/ticketDTO"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

type TicketUsecase interface {
	GetTicketById(id string) *ticketDTO.ReadTicketDTO
	GetAllTickets() []*ticketDTO.ReadTicketDTO
	CreateTicket(ticket ticketDTO.CreateTicketDTO) error
}

type ticketUsecase struct {
	ticketService service.TicketService
}

func NewTicketUsecase(ticketService service.TicketService) *ticketUsecase {
	return &ticketUsecase{ticketService: ticketService}
}

func (tUsecase *ticketUsecase) GetTicketById(id string) *ticketDTO.ReadTicketDTO {
	//TODO implement me
	panic("implement me")
}

func (tUsecase *ticketUsecase) GetAllTickets() []*ticketDTO.ReadTicketDTO {
	tickets := tUsecase.ticketService.GetAllTickets()

	readTicketsDTO := make([]*ticketDTO.ReadTicketDTO, 0, len(tickets))
	for _, ticket := range tickets {
		tempTicket := ticketDTO.ReadTicketDTO(ticket)
		readTicketsDTO = append(readTicketsDTO, &tempTicket)
	}

	return readTicketsDTO
}

func (tUsecase *ticketUsecase) CreateTicket(createTicket ticketDTO.CreateTicketDTO) error {
	id := uuid.New().String()
	ticket := ticketDTO.NewTicketFromCreateTicketDTO(id, createTicket)

	err := tUsecase.ticketService.CreateTicket(*ticket)
	if err != nil {
		log.Errorf("error creating ticket: %s", err.Error())
		return err
	}

	return nil
}
