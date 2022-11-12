package usecase

import (
	"api-app/internal/domain/entity"
	"api-app/internal/domain/service"
)

type UserUsecase interface {
	Usecase[entity.User, entity.UserView]
}

type userUsecase struct {
	service.UserService
}

var _ UserUsecase = (*userUsecase)(nil)

func NewUserUsecase(userService service.UserService) UserUsecase {
	return &userUsecase{UserService: userService}
}
