package service

import (
	"api-app/internal/domain/entity"
	"api-app/internal/domain/storage"
)

type FlightService interface {
	GetFlightById(id string) entity.Flight
	GetAllFlights() []entity.Flight
	GetAllFlightsMap() map[string]entity.Flight
	CreateFlight(f entity.Flight) error
	DeleteFlightById(id string) error
}

var _ Service = (*FlightService)(nil)

type flightService struct {
	storage storage.FlightStorage
}

var _ FlightService = (*flightService)(nil)

func (fService *flightService) GetFlightById(id string) entity.Flight {
	//TODO implement me
	panic("implement me")
}

func (fService *flightService) GetAllFlights() []entity.Flight {
	return fService.storage.GetAll()
}

func (fService *flightService) GetAllFlightsMap() map[string]entity.Flight {
	flightsMap := map[string]entity.Flight{}
	for _, flight := range fService.storage.GetAll() {
		flightsMap[flight.Id] = flight
	}

	return flightsMap
}

func (fService *flightService) CreateFlight(f entity.Flight) error {
	err := fService.storage.Store(f)
	if err != nil {
		return err
	}

	return nil
}

func (fService *flightService) DeleteFlightById(id string) error {
	//TODO implement me
	panic("implement me")
}

func NewFlightService(storage storage.FlightStorage) *flightService {
	return &flightService{storage: storage}
}
