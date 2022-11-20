package product

import (
	"api-app/internal/config"
	"api-app/internal/infrastructure/repository"
	"api-app/pkg/db/postgresql"
)

type Repositories interface {
	FlightRepo() repository.FlightRepository
	TicketRepo() repository.TicketRepository
	UserRepo() repository.UserRepository
}

type repos struct {
	flightRepo repository.FlightRepository
	ticketRepo repository.TicketRepository
	userRepo   repository.UserRepository
}

func NewRepositories(pgPool postgresql.PGXPool, dbCfg *config.PostgresConfig) (Repositories, error) {
	return &repos{
		flightRepo: repository.NewFlightRepository(pgPool, *dbCfg),
		ticketRepo: repository.NewTicketRepository(pgPool, *dbCfg),
		userRepo:   repository.NewUserRepository(pgPool, *dbCfg),
	}, nil
}

func (r *repos) FlightRepo() repository.FlightRepository {
	return r.flightRepo
}

func (r *repos) TicketRepo() repository.TicketRepository {
	return r.ticketRepo
}

func (r *repos) UserRepo() repository.UserRepository {
	return r.userRepo
}
