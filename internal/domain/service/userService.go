package service

import (
	"context"
	"github.com/achillescres/saina-api/internal/domain/entity"
	"github.com/achillescres/saina-api/internal/domain/storage"
	"github.com/achillescres/saina-api/internal/domain/storage/dto"
	"github.com/achillescres/saina-api/pkg/object/oid"
)

type UserService interface {
	PrimitiveService[entity.User, dto.UserCreate]
	GetByLoginAndHashedPassword(ctx context.Context, login, hashedPassword string) (*entity.User, error)
}

type userService struct {
	storage storage.UserStorage
}

var _ UserService = (*userService)(nil)

func (uS *userService) GetByLoginAndHashedPassword(ctx context.Context, login, hashedPassword string) (*entity.User, error) {
	return uS.storage.GetByLoginAndHashedPassword(ctx, login, hashedPassword)
}

func (uS *userService) GetById(ctx context.Context, id oid.Id) (*entity.User, error) {
	return uS.storage.GetById(ctx, id)
}

func (uS *userService) GetAll(ctx context.Context) ([]*entity.User, error) {
	return uS.GetAll(ctx)
}

func (uS *userService) GetAllByMap(ctx context.Context) (map[oid.Id]*entity.User, error) {
	return uS.GetAllByMap(ctx)
}

func (uS *userService) Store(ctx context.Context, uC dto.UserCreate) (*entity.User, error) {
	return uS.Store(ctx, uC)
}

func (uS *userService) DeleteById(ctx context.Context, id oid.Id) (*entity.User, error) {
	return uS.DeleteById(ctx, id)
}

func NewUserService(storage storage.UserStorage) UserService {
	return &userService{storage: storage}
}
