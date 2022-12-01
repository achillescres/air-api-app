package entity

import "api-app/pkg/object/oid"

type User struct {
	Entity         `json:"-" db:"-" binding:"-"`
	Id             oid.Id `json:"id" db:"id" binding:"required"`
	Login          string `json:"login" db:"login" binding:"required, max=30"`
	HashedPassword string `json:"hashedPassword" db:"hashed_password" binding:"required"`
}

func NewUser(id oid.Id, login string, hashedPassword string) *User {
	return &User{Id: id, Login: login, HashedPassword: hashedPassword}
}
