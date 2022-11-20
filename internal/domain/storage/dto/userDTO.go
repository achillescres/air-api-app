package dto

import "api-app/internal/domain/entity"

type UserCreate struct {
	Create
	Login    string `json:"login" `
	Password string `json:"password" `
}

func (uC *UserCreate) ToUserView(hashedPassword string) entity.UserView {
	return entity.UserView{
		Login:          uC.Login,
		HashedPassword: hashedPassword,
	}
}
