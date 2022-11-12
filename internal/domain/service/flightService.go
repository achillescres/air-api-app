package service

import (
	"api-app/internal/domain/entity"
	"api-app/internal/domain/storage"
	"api-app/pkg/object/oid"
)

type FlightService interface {
	Service[entity.Flight, entity.FlightView]
}

type flightService struct {
	storage storage.FlightStorage
}

var _ FlightService = (*flightService)(nil)

func (fService *flightService) GetById(id oid.Id) (entity.Flight, error) {
	//TODO implement me
	panic("implement me")
}

func (fService *flightService) GetAll() ([]entity.Flight, error) {
	return fService.storage.GetAll()
}

func (fService *flightService) GetAllByMap() (map[oid.Id]entity.Flight, error) {
	flightsMap := map[oid.Id]entity.Flight{}
	flights, err := fService.storage.GetAll()
	if err != nil {
		return nil, err
	}
	for _, flight := range flights {
		flightsMap[flight.Id] = flight
	}

	return flightsMap, nil
}

func (fService *flightService) Store(flightView entity.FlightView) (entity.Flight, error) {
	f, err := fService.storage.Store(flightView)
	return f, err
}

func (fService *flightService) DeleteById(id oid.Id) (entity.Flight, error) {
	return fService.storage.DeleteById(id)
}

func NewFlightService(storage storage.FlightStorage) FlightService {
	return &flightService{storage: storage}
}
