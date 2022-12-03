package service

import (
	"context"
	"github.com/achillescres/saina-api/internal/domain/entity"
	"github.com/achillescres/saina-api/internal/domain/storage/dto"
	"github.com/achillescres/saina-api/pkg/object/oid"
)

type PrimitiveService[Entity entity.Entity, Create dto.Create] interface {
	GetById(ctx context.Context, id oid.Id) (*Entity, error)
	GetAll(ctx context.Context) ([]*Entity, error)
	GetAllByMap(ctx context.Context) (map[oid.Id]*Entity, error)
	Store(ctx context.Context, c Create) (*Entity, error)
	DeleteById(ctx context.Context, id oid.Id) (*Entity, error)
}
