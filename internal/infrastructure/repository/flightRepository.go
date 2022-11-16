package repository

import (
	"api-app/internal/config"
	"api-app/internal/domain/entity"
	"api-app/internal/domain/storage"
	"api-app/internal/domain/storage/dto"
	"api-app/pkg/db/postgresql"
	"api-app/pkg/object/oid"
	"context"
)

type FlightRepository storage.Storage[entity.Flight, entity.FlightView, dto.FLightCreate]

type flightRepository struct {
	pool postgresql.Pool
	cfg  config.DBConfig
}

var _ FlightRepository = (*flightRepository)(nil)

func NewFlightRepository(pool postgresql.Pool, cfg config.DBConfig) FlightRepository {
	return &flightRepository{pool: pool, cfg: cfg}
}

func (fRepo *flightRepository) GetById(ctx context.Context, id oid.Id) (entity.Flight, error) {
	// TODO implement me
	panic("implement me")
}

func (fRepo *flightRepository) GetAll(ctx context.Context) ([]entity.Flight, error) {
	// TODO impl me
	panic("impl me")
}

func (fRepo *flightRepository) Store(ctx context.Context, fC dto.FLightCreate) (entity.Flight, error) {
	// TODO impl me
	panic("impl me")
}

func (fRepo *flightRepository) DeleteById(ctx context.Context, id oid.Id) (entity.Flight, error) {
	//TODO implement me
	panic("implement me")
}
