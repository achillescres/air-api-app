package service

import (
	"api-app/internal/domain/entity"
	"api-app/internal/domain/storage"
	"api-app/internal/domain/storage/dto"
	"api-app/pkg/object/oid"
	"context"
)

type FlightService interface {
	Service[entity.Flight, entity.FlightView, dto.FLightCreate]
}

type flightService struct {
	storage storage.FlightStorage
}

var _ FlightService = (*flightService)(nil)

func (fService *flightService) GetById(ctx context.Context, id oid.Id) (entity.Flight, error) {
	//TODO implement me
	panic("implement me")
}

func (fService *flightService) GetAll(ctx context.Context) ([]entity.Flight, error) {
	return fService.storage.GetAll(ctx)
}

func (fService *flightService) GetAllByMap(ctx context.Context) (map[oid.Id]entity.Flight, error) {
	flightsMap := map[oid.Id]entity.Flight{}
	flights, err := fService.storage.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	for _, flight := range flights {
		flightsMap[flight.Id] = flight
	}

	return flightsMap, nil
}

func (fService *flightService) Store(ctx context.Context, fC dto.FLightCreate) (entity.Flight, error) {
	f, err := fService.storage.Store(ctx, fC)
	return f, err
}

func (fService *flightService) DeleteById(ctx context.Context, id oid.Id) (entity.Flight, error) {
	return fService.storage.DeleteById(ctx, id)
}

func NewFlightService(storage storage.FlightStorage) FlightService {
	return &flightService{storage: storage}
}
