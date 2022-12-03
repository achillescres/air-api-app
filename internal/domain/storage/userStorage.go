package storage

import (
	"context"
	"github.com/achillescres/saina-api/internal/domain/entity"
	"github.com/achillescres/saina-api/internal/domain/storage/dto"
)

type UserStorage interface {
	Storage[entity.User, dto.UserCreate]
	GetByLoginAndHashedPassword(ctx context.Context, login, hashedPassword string) (*entity.User, error)
}
