package service

import (
	"api-app/internal/domain/entity"
)

type Service[Entity entity.Entity, View entity.View] interface {
	GetById(id id.Id) (Entity, error)
	GetAll() ([]Entity, error)
	GetAllByMap() (map[id.Id]Entity, error)
	Store(v View) (Entity, error)
	DeleteById(id id.Id) (Entity, error)
}
