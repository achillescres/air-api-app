package composite

import (
	"api-app/internal/adapter/service"
	"api-app/internal/repository"
)

func GenerateFlightComposite() service.FlightService {
	repo := repository.NewFlightRepository()
	serv := service.NewFlightService(repo)
	return serv
}
