package entity

import "api-app/pkg/object/oid"

type User struct {
	Entity `json:"-" db:"-"`
	Id     oid.Id `json:"id" db:"id"`

	// View
	Login          string `json:"login" db:"login"`
	HashedPassword string `json:"hashedPassword" db:"hashed_password"`
}

func (u *User) ToView() *UserView {
	return &UserView{
		Login:          u.Login,
		HashedPassword: u.HashedPassword,
	}
}

type UserView struct {
	View           `json:"-" db:"-"`
	Login          string `json:"login" db:"login"`
	HashedPassword string `json:"hashedPassword" db:"hashed_password"`
}

func (uV UserView) ToEntity(id oid.Id) *User {
	return &User{
		Id:             id,
		Login:          uV.Login,
		HashedPassword: uV.HashedPassword,
	}
}
