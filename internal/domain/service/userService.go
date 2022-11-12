package service

import (
	"api-app/internal/domain/entity"
	"api-app/internal/domain/storage"
	"api-app/pkg/object/oid"
)

type UserService interface {
	Service[entity.User, entity.UserView]
}

type userService struct {
	storage storage.UserStorage
}

var _ UserService = (*userService)(nil)

func (uS *userService) GetById(id oid.Id) (entity.User, error) {
	return uS.GetById(id)
}

func (uS *userService) GetAll() ([]entity.User, error) {
	return uS.GetAll()
}

func (uS *userService) GetAllByMap() (map[oid.Id]entity.User, error) {
	return uS.GetAllByMap()
}

func (uS *userService) Store(uV entity.UserView) (entity.User, error) {
	return uS.Store(uV)
}

func (uS *userService) DeleteById(id oid.Id) (entity.User, error) {
	return uS.DeleteById(id)
}

func NewUserService(storage storage.UserStorage) UserService {
	return &userService{storage: storage}
}
