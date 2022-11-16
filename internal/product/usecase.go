package product

import (
	"api-app/internal/config"
	"api-app/internal/domain/usecase"
)

type Usecase interface {
	FlightUc() usecase.FlightUsecase
	TicketUc() usecase.TicketUsecase
	UserUc() usecase.UserUsecase
}

type uc struct {
	flightUc usecase.FlightUsecase
	ticketUc usecase.TicketUsecase
	userUc   usecase.UserUsecase
}

func (u *uc) FlightUc() usecase.FlightUsecase {
	return u.flightUc
}

func (u *uc) TicketUc() usecase.TicketUsecase {
	return u.ticketUc
}

func (u *uc) UserUc() usecase.UserUsecase {
	return u.userUc
}

func NewUsecase(service Service, cfg *config.UsecaseConfig) (Usecase, error) {
	return &uc{
		flightUc: usecase.NewFlightUsecase(service.FlightService(), service.TicketService(), *cfg),
		ticketUc: usecase.NewTicketUsecase(service.TicketService()),
		userUc:   usecase.NewUserUsecase(service.UserService()),
	}, nil
}
