package usecase

import (
	"api-app/internal/domain/entity"
	"api-app/internal/domain/service"
	"api-app/internal/domain/storage/dto"
	"context"
)

type UserUsecase interface {
	Usecase
	StoreUser(ctx context.Context, uC dto.UserCreate) (*entity.User, error)
}

type userUsecase struct {
	userService service.UserService
}

var _ UserUsecase = (*userUsecase)(nil)

func NewUserUsecase(userService service.UserService) UserUsecase {
	return &userUsecase{userService: userService}
}

func (uU *userUsecase) StoreUser(ctx context.Context, uC dto.UserCreate) (*entity.User, error) {
	store, err := uU.userService.Store(ctx, uC)
	if err != nil {
		return nil, err
	}
	return store, nil
}
