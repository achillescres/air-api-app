package storage

import (
	"context"
	"github.com/achillescres/saina-api/internal/domain/entity"
	"github.com/achillescres/saina-api/internal/domain/storage/dto"
)

type RefreshTokenStorage interface {
	Storage[entity.RefreshToken, dto.RefreshTokenCreate]
	GetByToken(ctx context.Context, refreshToken string) (*entity.RefreshToken, error)
}
