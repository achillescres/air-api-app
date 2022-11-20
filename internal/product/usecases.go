package product

import (
	"api-app/internal/config"
	"api-app/internal/domain/usecase"
)

type Usecases interface {
	FlightUc() usecase.FlightUsecase
	TicketUc() usecase.TicketUsecase
	UserUc() usecase.UserUsecase
}

type ucs struct {
	flightUc usecase.FlightUsecase
	ticketUc usecase.TicketUsecase
	userUc   usecase.UserUsecase
}

func (u *ucs) FlightUc() usecase.FlightUsecase {
	return u.flightUc
}

func (u *ucs) TicketUc() usecase.TicketUsecase {
	return u.ticketUc
}

func (u *ucs) UserUc() usecase.UserUsecase {
	return u.userUc
}

func NewUsecases(service Services, cfg *config.UsecaseConfig) (Usecases, error) {
	return &ucs{
		flightUc: usecase.NewFlightUsecase(service.FlightService(), service.TicketService(), *cfg),
		ticketUc: usecase.NewTicketUsecase(service.TicketService()),
		userUc:   usecase.NewUserUsecase(service.UserService()),
	}, nil
}
