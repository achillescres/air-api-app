package service

import (
	"api-app/internal/domain/entity"
	"api-app/internal/domain/storage"
	"github.com/google/uuid"
)

type TicketService interface {
	GetTicketById(id string) entity.Ticket
	GetAllTickets() []entity.Ticket
	GetAllTicketsMap() map[string]entity.Ticket
	CreateTicket(fV entity.TicketView) error
	DeleteTicketById(id string) error
}

var _ Service = (*TicketService)(nil)

type ticketService struct {
	storage storage.TicketStorage
}

var _ TicketService = (*ticketService)(nil)

func (tService *ticketService) GetAllTicketsMap() map[string]entity.Ticket {
	//TODO implement me
	panic("implement me")
}

func (tService *ticketService) GetTicketById(id string) entity.Ticket {
	//TODO implement me
	panic("implement me")
}

func (tService *ticketService) GetAllTickets() []entity.Ticket {
	return tService.storage.GetAll()
}

func (tService *ticketService) CreateTicket(fV entity.TicketView) error {
	id := uuid.New().String()
	err := tService.storage.Store(*entity.FromTicketView(id, fV))
	if err != nil {
		return err
	}

	return nil
}

func (tService *ticketService) DeleteTicketById(id string) error {
	//TODO implement me
	panic("implement me")
}

func NewTicketService(storage storage.TicketStorage) *ticketService {
	return &ticketService{storage: storage}
}
