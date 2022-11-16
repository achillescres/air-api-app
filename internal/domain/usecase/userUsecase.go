package usecase

import (
	"api-app/internal/domain/entity"
	"api-app/internal/domain/service"
	"api-app/internal/domain/storage/dto"
)

type UserUsecase interface {
	Usecase[entity.User, entity.UserView, dto.UserCreate]
}

type userUsecase struct {
	service.UserService
}

var _ UserUsecase = (*userUsecase)(nil)

func NewUserUsecase(userService service.UserService) UserUsecase {
	return &userUsecase{UserService: userService}
}
