package dto

import (
	"github.com/achillescres/saina-api/internal/domain/entity"
	"github.com/achillescres/saina-api/pkg/object/oid"
)

type UserCreate struct {
	Create   `json:"-" db:"-" binding:"-"`
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

var _ Create = (*UserCreate)(nil)

func NewUserCreate(login string, password string) *UserCreate {
	return &UserCreate{Login: login, Password: password}
}

func (uC *UserCreate) ToEntity(id oid.Id, hashedPassword string) *entity.User {
	return entity.NewUser(id, uC.Login, hashedPassword)
}
