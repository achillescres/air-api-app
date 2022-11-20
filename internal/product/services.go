package product

import "api-app/internal/domain/service"

type Services interface {
	FlightService() service.FlightService
	TicketService() service.TicketService
	UserService() service.UserService
}

type servs struct {
	flightService service.FlightService
	ticketService service.TicketService
	userService   service.UserService
}

func (s *servs) FlightService() service.FlightService {
	return s.flightService
}

func (s *servs) TicketService() service.TicketService {
	return s.ticketService
}

func (s *servs) UserService() service.UserService {
	return s.userService
}

func NewServices(
	repo Repositories,
) (Services, error) {
	return &servs{
		flightService: service.NewFlightService(repo.FlightRepo()),
		ticketService: service.NewTicketService(repo.TicketRepo()),
		userService:   service.NewUserService(repo.UserRepo()),
	}, nil
}
