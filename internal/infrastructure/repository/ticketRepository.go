package repository

import (
	"context"
	"errors"
	"github.com/achillescres/saina-api/internal/domain/entity"
	"github.com/achillescres/saina-api/internal/domain/storage"
	"github.com/achillescres/saina-api/internal/domain/storage/dto"
	"github.com/achillescres/saina-api/pkg/db/postgresql"
	"github.com/achillescres/saina-api/pkg/object/oid"
	"github.com/georgysavva/scany/v2/pgxscan"
	log "github.com/sirupsen/logrus"
)

type TicketRepository interface {
	storage.TicketStorage
}

type ticketRepository struct {
	pool postgresql.PGXPool
}

var _ storage.TicketStorage = (*ticketRepository)(nil)

func NewTicketRepository(pool postgresql.PGXPool) TicketRepository {
	return &ticketRepository{pool: pool}
}

func (tRepo *ticketRepository) GetById(ctx context.Context, id oid.Id) (*entity.Ticket, error) {
	// TODO implement me
	panic("implement me")
}

func (tRepo *ticketRepository) GetAll(ctx context.Context) ([]*entity.Ticket, error) {
	rows, err := tRepo.pool.Query(ctx, "SELECT * FROM public.tickets")
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	var tickets []*entity.Ticket
	err = pgxscan.ScanAll(&tickets, rows)
	if err != nil {
		log.Errorf("error scanning rows of tickets: %s", err.Error())
		return nil, err
	}

	return tickets, err
}

func (tRepo *ticketRepository) GetAllInMap(ctx context.Context) (map[oid.Id]*entity.Ticket, error) {
	//TODO implement me
	panic("implement me")
}

func (tRepo *ticketRepository) GetAllByFlightId(ctx context.Context, flightId oid.Id) ([]*entity.Ticket, error) {
	// TODO implement me
	panic("implement me")
}

func (tRepo *ticketRepository) Store(ctx context.Context, tC dto.TicketCreate) (*entity.Ticket, error) {
	query, err := tRepo.pool.Query(
		ctx,
		"INSERT INTO public.tickets (flight_id, airl_code, flt_num, flt_date, ticket_code, ticket_capacity, ticket_type, amount, total_cash, correctly_parsed) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING (id)",
		tC.FlightId,
		tC.AirlCode,
		tC.FltNum,
		tC.FltDate,
		tC.TicketCode,
		tC.TicketCapacity,
		tC.TicketType,
		tC.Amount,
		tC.TotalCash,
		tC.CorrectlyParsed,
	)
	defer query.Close()
	if err != nil {
		log.Errorf("error inserting new flight: %s\n", err.Error())
		return nil, err
	}

	if !query.Next() {
		err := errors.New("error sql didn't return id of new Ticket")
		log.Errorln(err.Error())

		return nil, err // TODO WHAT TO DO WTF???!!!?
	}

	var id string
	err = query.Scan(&id)
	if err != nil {
		log.Errorf("error scanning new newTicket id: %s\n", err.Error())
		return nil, err // TODO WHAT TO DO WTF??!?!?!?!?
	}

	return tC.ToEntity(oid.ToId(id)), err
}

func (tRepo *ticketRepository) DeleteById(ctx context.Context, id oid.Id) (*entity.Ticket, error) {
	// TODO implement me
	panic("implement me")
}
