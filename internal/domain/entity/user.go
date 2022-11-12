package entity

type User struct {
	Id   id.Id `json:"id"`
	View UserView
}

var _ Entity = (*User)(nil)

type UserView struct {
	Login          string `json:"login"`
	HashedPassword string `json:"hashedPassword"`
}

func ToUserView(u User) UserView {
	return u.View
}

func FromUserView(id id.Id, uV UserView) *User {
	return &User{
		Id:   id,
		View: uV,
	}
}
