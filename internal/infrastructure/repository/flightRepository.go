package repository

import (
	"api-app/internal/domain/entity"
	"api-app/internal/domain/storage"
	"api-app/pkg/object/oid"
)

type FLightRepository storage.Storage[entity.Flight, entity.FlightView]

type flightRepository struct {
	collection map[oid.Id]entity.Flight
}

var _ FLightRepository = (*flightRepository)(nil)

func (fRepo *flightRepository) GetById(id oid.Id) (entity.Flight, error) {
	//TODO implement me
	panic("implement me")
}

func (fRepo *flightRepository) GetAll() ([]entity.Flight, error) {
	flights := make([]entity.Flight, 0, len(fRepo.collection))
	for _, flight := range fRepo.collection {
		flights = append(flights, flight)
	}

	return flights, nil
}

func (fRepo *flightRepository) Store(f entity.FlightView) (entity.Flight, error) {
	id := oid.NewId()
	newFlight := entity.FromFlightView(id, f)
	fRepo.collection[id] = *newFlight
	return *newFlight, nil
}

func (fRepo *flightRepository) DeleteById(id oid.Id) (entity.Flight, error) {
	//TODO implement me
	panic("implement me")
}

func NewFlightRepository() FLightRepository {
	return &flightRepository{collection: make(map[oid.Id]entity.Flight)}
}
