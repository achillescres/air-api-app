package repository

import (
	"api-app/internal/domain/entity"
	"api-app/internal/domain/storage"
	"api-app/pkg/object/oid"
)

type UserRepository interface {
	storage.Storage[entity.User, entity.UserView]
}

type userRepository struct {
	collection map[oid.Id]entity.User
}

var _ UserRepository = (*userRepository)(nil)

func (uR *userRepository) GetById(id oid.Id) (entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (uR *userRepository) GetAll() ([]entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (uR *userRepository) Store(uV entity.UserView) (entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (uR *userRepository) DeleteById(id oid.Id) (entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func NewUserRepository() UserRepository {
	return &userRepository{collection: map[oid.Id]entity.User{}}
}
