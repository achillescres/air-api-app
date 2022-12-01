package product

import (
	"api-app/internal/infrastructure/repository"
	"api-app/pkg/db/postgresql"
	"api-app/pkg/security/passlib"
)

type Repositories struct {
	FlightRepo       repository.FlightRepository
	TicketRepo       repository.TicketRepository
	UserRepo         repository.UserRepository
	RefreshTokenRepo repository.RefreshTokenRepository
}

func NewRepositories(pgPool postgresql.PGXPool, hashManager passlib.HashManager) (*Repositories, error) {
	return &Repositories{
		FlightRepo:       repository.NewFlightRepository(pgPool),
		TicketRepo:       repository.NewTicketRepository(pgPool),
		UserRepo:         repository.NewUserRepository(pgPool, hashManager),
		RefreshTokenRepo: repository.NewRefreshTokenRepository(pgPool),
	}, nil
}
