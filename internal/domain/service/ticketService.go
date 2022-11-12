package service

import (
	"api-app/internal/domain/entity"
	"api-app/internal/domain/storage"
	"api-app/pkg/object/oid"
)

type TicketService interface {
	Service[entity.Ticket, entity.TicketView]
}

type ticketService struct {
	storage storage.TicketStorage
}

var _ TicketService = (*ticketService)(nil)

func (tService *ticketService) GetAllByMap() (map[oid.Id]entity.Ticket, error) {
	ticketsMap := map[oid.Id]entity.Ticket{}
	tickets, err := tService.storage.GetAll()
	if err != nil {
		return nil, err
	}
	for _, ticket := range tickets {
		ticketsMap[ticket.Id] = ticket
	}

	return ticketsMap, nil
}

func (tService *ticketService) GetById(id oid.Id) (entity.Ticket, error) {
	return tService.GetById(id)
}

func (tService *ticketService) GetAll() ([]entity.Ticket, error) {
	return tService.storage.GetAll()
}

func (tService *ticketService) Store(tV entity.TicketView) (entity.Ticket, error) {
	return tService.storage.Store(tV)
}

func (tService *ticketService) DeleteById(id oid.Id) (entity.Ticket, error) {
	return tService.storage.DeleteById(id)
}

func NewTicketService(storage storage.TicketStorage) TicketService {
	return &ticketService{storage: storage}
}
