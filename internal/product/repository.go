package product

import (
	"api-app/internal/config"
	"api-app/internal/infrastructure/repository"
	"api-app/pkg/db/postgresql"
)

type Repository interface {
	FlightRepo() repository.FlightRepository
	TicketRepo() repository.TicketRepository
	UserRepo() repository.UserRepository
}

type repo struct {
	flightRepo repository.FlightRepository
	ticketRepo repository.TicketRepository
	userRepo   repository.UserRepository
}

func (r *repo) FlightRepo() repository.FlightRepository {
	return r.flightRepo
}

func (r *repo) TicketRepo() repository.TicketRepository {
	return r.ticketRepo
}

func (r *repo) UserRepo() repository.UserRepository {
	return r.userRepo
}

func NewRepository(pgPool postgresql.Pool, dbCfg *config.DBConfig) (Repository, error) {
	return &repo{
		flightRepo: repository.NewFlightRepository(pgPool, *dbCfg),
		ticketRepo: repository.NewTicketRepository(pgPool, *dbCfg),
		userRepo:   repository.NewUserRepository(pgPool, *dbCfg),
	}, nil
}
