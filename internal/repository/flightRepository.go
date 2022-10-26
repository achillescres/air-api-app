package repository

import (
	"api-app/internal/adapter/storage"
	"api-app/internal/entity"
)

type flightRepository struct {
	collection map[string]*entity.Flight
}

var _ Repository = (*flightRepository)(nil)
var _ storage.FlightStorage = (*flightRepository)(nil)

func (fRepo *flightRepository) GetById(id string) *entity.Flight {
	//TODO implement me
	panic("implement me")
}

func (fRepo *flightRepository) GetAll() []*entity.Flight {
	//TODO implement me
	flights := make([]*entity.Flight, 0, len(fRepo.collection))
	for _, flight := range fRepo.collection {
		flights = append(flights, flight)
	}

	return flights
}

func (fRepo *flightRepository) Store(f entity.Flight) error {
	//TODO implement me
	panic("implement me")
}

func (fRepo *flightRepository) DeleteById(id string) (*entity.Flight, error) {
	//TODO implement me
	panic("implement me")
}

func NewFlightRepository() *flightRepository {
	return &flightRepository{}
}
