package service

import (
	"api-app/internal/domain/dto"
	"api-app/internal/domain/entity"
	"api-app/pkg/object/oid"
	"context"
)

type PrimitiveService[Entity entity.Entity, View entity.View, Create dto.Create] interface {
	GetById(ctx context.Context, id oid.Id) (*Entity, error)
	GetAll(ctx context.Context) ([]*Entity, error)
	GetAllByMap(ctx context.Context) (map[oid.Id]*Entity, error)
	Store(ctx context.Context, c Create) (*Entity, error)
	DeleteById(ctx context.Context, id oid.Id) (*Entity, error)
}
