package httpHandler

import (
	"api-app/internal/domain/entity"
	"api-app/internal/domain/usecase"
	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	Handler[entity.Flight]
	Register(c *gin.Context)
	Login(c *gin.Context)
	Logout(c *gin.Context)
}

type userHandler struct {
	usecase.UserUsecase
}

func (uH *userHandler) Register(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (uH *userHandler) Login(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (uH *userHandler) Logout(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func NewUserHandler(userUsecase usecase.UserUsecase) UserHandler {
	return &userHandler{UserUsecase: userUsecase}
}
