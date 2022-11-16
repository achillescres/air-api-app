package dto

import "api-app/internal/domain/entity"

type UserCreate struct {
	Create
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (uC *UserCreate) ToUserView(hashedPassword string) entity.UserView {
	return entity.UserView{
		Login:          uC.Login,
		HashedPassword: hashedPassword,
	}
}
