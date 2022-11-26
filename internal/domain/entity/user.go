package entity

import "api-app/pkg/object/oid"

type User struct {
	Entity `json:"-" db:"-" binding:"-"`
	Id     oid.Id `json:"id" db:"id" binding:"required"`

	// View
	Login          string `json:"login" db:"login" binding:"required, max=30"`
	HashedPassword string `json:"hashedPassword" db:"hashed_password" binding:"required"`
}

func (u *User) ToView() *UserView {
	return &UserView{
		Login:          u.Login,
		HashedPassword: u.HashedPassword,
	}
}

type UserView struct {
	View           `json:"-" db:"-" binding:"-"`
	Login          string `json:"login" db:"login" binding:"required, max=30"`
	HashedPassword string `json:"hashedPassword" db:"hashed_password" binding:"required"`
}

func (uV UserView) ToEntity(id oid.Id) *User {
	return &User{
		Id:             id,
		Login:          uV.Login,
		HashedPassword: uV.HashedPassword,
	}
}
