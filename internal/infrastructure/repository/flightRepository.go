package repository

import (
	"api-app/internal/domain/entity"
	"api-app/internal/domain/storage"
	"errors"
	"github.com/google/uuid"
)

type flightRepository struct {
	collection map[string]entity.Flight
}

var _ Repository = (*flightRepository)(nil)
var _ storage.FlightStorage = (*flightRepository)(nil)

func (fRepo *flightRepository) GetById(id string) entity.Flight {
	//TODO implement me
	panic("implement me")
}

func (fRepo *flightRepository) GetAll() []entity.Flight {
	//TODO implement me
	flights := make([]entity.Flight, 0, len(fRepo.collection))
	for _, flight := range fRepo.collection {
		flights = append(flights, flight)
	}

	return flights
}

func (fRepo *flightRepository) Store(f entity.FlightView) (entity.Flight, error) {
	id := uuid.New().String()
	_, contains := fRepo.collection[id]
	if !contains {
		newFlight := entity.FullFlight()
		fRepo.collection[id] = f
	} else {
		return errors.New("error already contains this id")
	}

	return nil
}

func (fRepo *flightRepository) DeleteById(id string) (entity.Flight, error) {
	//TODO implement me
	panic("implement me")
}

func NewFlightRepository() *flightRepository {
	return &flightRepository{collection: make(map[string]entity.Flight)}
}
