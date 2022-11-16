package product

import "api-app/internal/domain/service"

type Service interface {
	FlightService() service.FlightService
	TicketService() service.TicketService
	UserService() service.UserService
}

type serv struct {
	flightService service.FlightService
	ticketService service.TicketService
	userService   service.UserService
}

func (s *serv) FlightService() service.FlightService {
	return s.flightService
}

func (s *serv) TicketService() service.TicketService {
	return s.ticketService
}

func (s *serv) UserService() service.UserService {
	return s.userService
}

func NewService(
	repo Repository,
) (Service, error) {
	return &serv{
		flightService: service.NewFlightService(repo.FlightRepo()),
		ticketService: service.NewTicketService(repo.TicketRepo()),
		userService:   service.NewUserService(repo.UserRepo()),
	}, nil
}
