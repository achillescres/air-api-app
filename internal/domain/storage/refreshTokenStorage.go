package storage

import (
	"api-app/internal/domain/entity"
	"api-app/internal/domain/storage/dto"
	"context"
)

type RefreshTokenStorage interface {
	Storage[entity.RefreshToken, dto.RefreshTokenCreate]
	GetByToken(ctx context.Context, refreshToken string) (*entity.RefreshToken, error)
}
