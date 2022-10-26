package storage

import "api-app/internal/entity"

type FlightStorage interface {
	GetById(id string) *entity.Flight
	GetAll() []*entity.Flight

	Store(f entity.Flight) error
	DeleteById(id string) (*entity.Flight, error)
}

var _ Storage = (*FlightStorage)(nil)
