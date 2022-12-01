package storage

import (
	"api-app/internal/domain/entity"
	"api-app/internal/domain/storage/dto"
	"api-app/pkg/object/oid"
	"context"
)

type Storage[Entity entity.Entity, Create dto.Create] interface {
	GetById(ctx context.Context, id oid.Id) (*Entity, error)
	GetAll(ctx context.Context) ([]*Entity, error)
	GetAllInMap(ctx context.Context) (map[oid.Id]*Entity, error)
	Store(ctx context.Context, f Create) (*Entity, error)
	DeleteById(ctx context.Context, id oid.Id) (*Entity, error)
}
