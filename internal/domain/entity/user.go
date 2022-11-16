package entity

import "api-app/pkg/object/oid"

type User struct {
	Entity
	Id   oid.Id   `json:"id" binding:"required"`
	View UserView `json:"view" binding:"required"`
}

type UserView struct {
	View
	Login          string `json:"login" binding:"required"`
	HashedPassword string `json:"hashedPassword" binding:"required"`
}

func ToUserView(u User) UserView {
	return u.View
}

func FromUserView(id oid.Id, uV UserView) *User {
	return &User{
		Id:   id,
		View: uV,
	}
}
