package composite

import (
	"api-app/internal/domain/service"
	"api-app/internal/repository"
)

func GenerateFlightComposite() service.FlightService {
	repo := repository.NewFlightRepository()
	serv := service.NewFlightService(repo)
	return serv
}
