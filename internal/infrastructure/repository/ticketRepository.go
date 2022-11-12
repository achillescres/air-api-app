package repository

import (
	"api-app/internal/domain/entity"
	"api-app/internal/domain/storage"
	"api-app/pkg/object/oid"
)

type TicketRepository interface {
	storage.Storage[entity.Ticket, entity.TicketView]
}

type ticketRepository struct {
	collection map[oid.Id]entity.Ticket
}

var _ storage.TicketStorage = (*ticketRepository)(nil)

func (tRepo *ticketRepository) GetById(id oid.Id) (entity.Ticket, error) {
	//TODO implement me
	panic("implement me")
}

func (tRepo *ticketRepository) GetAll() ([]entity.Ticket, error) {
	tickets := make([]entity.Ticket, 0, len(tRepo.collection))
	for _, ticket := range tRepo.collection {
		tickets = append(tickets, ticket)
	}

	return tickets, nil
}

func (tRepo *ticketRepository) GetAllByFlightId(flightId oid.Id) ([]entity.Ticket, error) {
	//TODO implement me
	panic("implement me")
}

func (tRepo *ticketRepository) Store(t entity.TicketView) (entity.Ticket, error) {
	id := oid.NewId()
	_, contains := tRepo.collection[id]
	if !contains {
		tRepo.collection[id] = *entity.FromTicketView(id, t)
	}

	return tRepo.collection[id], nil
}

func (tRepo *ticketRepository) DeleteById(id oid.Id) (entity.Ticket, error) {
	//TODO implement me
	panic("implement me")
}

func NewTicketRepository() TicketRepository {
	return &ticketRepository{collection: make(map[oid.Id]entity.Ticket)}
}
