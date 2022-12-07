package product

import (
	repository2 "github.com/achillescres/saina-api/internal/infrastructure/repository"
	"github.com/achillescres/saina-api/pkg/db/postgresql"
	"github.com/achillescres/saina-api/pkg/security/passlib"
)

type Repositories struct {
	FlightRepo       repository2.FlightRepository
	TicketRepo       repository2.TicketRepository
	UserRepo         repository2.UserRepository
	RefreshTokenRepo repository2.RefreshTokenRepository
}

func NewRepositories(pgPool postgresql.PGXPool, hashManager passlib.HashManager) (*Repositories, error) {
	return &Repositories{
		FlightRepo:       repository2.NewFlightRepository(pgPool),
		TicketRepo:       repository2.NewTicketRepository(pgPool),
		UserRepo:         repository2.NewUserRepository(pgPool, hashManager),
		RefreshTokenRepo: repository2.NewRefreshTokenRepository(pgPool),
	}, nil
}
