package storage

import (
	"api-app/internal/domain/entity"
	"api-app/pkg/object/oid"
)

type Storage[Entity entity.Entity, View entity.View] interface {
	GetById(id oid.Id) (Entity, error)
	GetAll() ([]Entity, error)
	Store(f View) (Entity, error)
	DeleteById(id oid.Id) (Entity, error)
}
