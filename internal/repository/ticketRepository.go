package repository

import (
	"api-app/internal/domain/entity"
	"api-app/internal/domain/storage"
)

type ticketRepository struct {
	collection map[string]*entity.Ticket
}

var _ Repository = (*ticketRepository)(nil)
var _ storage.TicketStorage = (*ticketRepository)(nil)

func (tRepo *ticketRepository) GetById(id string) *entity.Ticket {
	//TODO implement me
	panic("implement me")
}

func (tRepo *ticketRepository) GetAll() []*entity.Ticket {
	tickets := make([]*entity.Ticket, 0, len(tRepo.collection))
	for _, ticket := range tRepo.collection {
		tickets = append(tickets, ticket)
	}

	return tickets
}

func (tRepo *ticketRepository) GetAllByFlightId(flightId string) []*entity.Ticket {
	//TODO implement me
	panic("implement me")
}

func (tRepo *ticketRepository) Store(t entity.Ticket) error {
	_, contains := tRepo.collection[t.Id]
	if !contains {
		tRepo.collection[t.Id] = &t
	}

	return nil
}

func (tRepo *ticketRepository) DeleteById(id string) (*entity.Ticket, error) {
	//TODO implement me
	panic("implement me")
}

func NewTicketRepository() *ticketRepository {
	return &ticketRepository{collection: make(map[string]*entity.Ticket)}
}
