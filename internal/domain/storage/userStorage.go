package storage

import (
	"api-app/internal/domain/entity"
	"api-app/internal/domain/storage/dto"
	"context"
)

type UserStorage interface {
	Storage[entity.User, dto.UserCreate]
	GetByLoginAndHashedPassword(ctx context.Context, login, hashedPassword string) (*entity.User, error)
}
