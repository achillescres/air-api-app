package httpHandler

import (
	"api-app/internal/domain/dto"
	"api-app/pkg/gin/ginresponse"
	"api-app/pkg/security/ajwt"
	"api-app/pkg/security/passlib"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Register is POST method, that registers users
// It gets the JSON with dto.UserCreate model
func (h *handler) Register(c *gin.Context) {
	registerInput := dto.RegisterInput{}
	err := c.ShouldBindJSON(&registerInput)
	if err != nil {
		ginresponse.WithError(c, http.StatusUnprocessableEntity, err, "invalid object")
		return
	}

	createUser := dto.UserCreate{
		Create:   nil,
		Login:    registerInput.Login,
		Password: registerInput.Password,
	}

	user, err := h.userService.Store(c, createUser)
	if err != nil {
		ginresponse.WithError(c, http.StatusInternalServerError, err, "couldn't store user")
		return
	}

	c.JSON(http.StatusOK, map[string]any{"id": user.Id})
}

func (h *handler) Login(c *gin.Context) {
	loginInput := dto.LoginInput{}
	err := c.ShouldBindJSON(&loginInput)
	if err != nil {
		ginresponse.WithError(c, http.StatusUnprocessableEntity, err, "invalid object")
		return
	}

	hashedPassword, err := passlib.Hash(loginInput.Password)
	user, err := h.userService.GetByLoginAndHashedPassword(c, loginInput.Login, hashedPassword)
	if err != nil {
		ginresponse.WithError(c, http.StatusBadRequest, err, "invalid login or password")
		return
	}

	token, err := ajwt.User(user.Login, user.HashedPassword)
	if err != nil {
		ginresponse.WithError(c, http.StatusInternalServerError, err, "couldn't generate token")
		return
	}

	h.jwtService.Store()
	h.refreshTokenService.Store()

	c.JSON(http.StatusOK, gin.H{"Bearer": token})
}

func (h *handler) Logout(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (h *handler) registerUser(r *gin.RouterGroup) {
	r.POST("/register", h.Register)
	r.POST("/login", h.Login)
	r.POST("/logout", h.Logout)
}
