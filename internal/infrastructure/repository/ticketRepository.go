package repository

import (
	"api-app/internal/config"
	"api-app/internal/domain/dto"
	"api-app/internal/domain/entity"
	"api-app/internal/domain/storage"
	"api-app/pkg/db/postgresql"
	"api-app/pkg/object/oid"
	"context"
	"errors"
	"github.com/georgysavva/scany/v2/pgxscan"
	log "github.com/sirupsen/logrus"
)

type TicketRepository interface {
	storage.Storage[entity.Ticket, entity.TicketView, dto.TicketCreate]
	GetAllByFlightId(ctx context.Context, flightId oid.Id) ([]*entity.Ticket, error)
}

type ticketRepository struct {
	pool postgresql.PGXPool
	cfg  config.PostgresConfig
}

var _ storage.TicketStorage = (*ticketRepository)(nil)

func NewTicketRepository(pool postgresql.PGXPool, cfg config.PostgresConfig) TicketRepository {
	return &ticketRepository{pool: pool, cfg: cfg}
}

func (tRepo *ticketRepository) GetById(ctx context.Context, id oid.Id) (entity.Ticket, error) {
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
		return &entity.Ticket{}, err
	}

	newTicket := tC.ToTicketView().ToEntity(oid.Undefined)
	if !query.Next() {
		err := errors.New("error there's no returned id from sql")
		log.Errorln(err.Error())
		return newTicket, err // TODO WHAT TO DO WTF???!!!?
	}
	var id string
	err = query.Scan(&id)
	if err != nil {
		log.Errorf("error scanning new newTicket id: %s\n", err.Error())
		return newTicket, err // TODO WHAT TO DO WTF??!?!?!?!?
	}

	newTicket.Id = oid.ToId(id)
	return newTicket, err
}

func (tRepo *ticketRepository) DeleteById(ctx context.Context, id oid.Id) (entity.Ticket, error) {
	// TODO implement me
	panic("implement me")
}
