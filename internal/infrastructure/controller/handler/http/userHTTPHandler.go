package httpHandler

import (
	"api-app/internal/domain/entity"
	"api-app/internal/domain/storage/dto"
	"api-app/internal/domain/usecase"
	"api-app/pkg/gin/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandler interface {
	Handler[entity.Flight]
	Register(c *gin.RouterGroup)
	Login(c *gin.Context)
	Logout(c *gin.Context)
}

type userHandler struct {
	usecase.UserUsecase
}

// Register is POST method, registers users
// get the JSON with dto.UserCreate model
func (uH *userHandler) Register(c *gin.Context) {
	createUser := dto.UserCreate{}
	err := c.BindJSON(&createUser)
	if err != nil {
		response.Error(c, http.StatusUnprocessableEntity, "invalid object format")
	}

	_, err = uH.Store(c, createUser)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "couldn't store user in repository")
	}
	c.JSON(http.StatusOK, createUser) // TODO response with createUser is not safe
}

func (uH *userHandler) Login(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (uH *userHandler) Logout(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (uH *userHandler) RegisterRouter(r *gin.RouterGroup) {
	r = r.Group("/user")
	r.POST("/register", uH.Register)
	r.POST("/login", uH.Login)
	r.POST("/logout", uH.Logout)
}

func NewUserHandler(userUsecase usecase.UserUsecase) UserHandler {
	return &userHandler{UserUsecase: userUsecase}
}
