package httpHandler

import (
	"errors"
	"github.com/achillescres/saina-api/internal/domain/service/sto"
	"github.com/achillescres/saina-api/internal/domain/storage"
	"github.com/achillescres/saina-api/pkg/gin/ginresponse"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Register is POST method, that registers users
// It gets the JSON with dto.UserCreate model
func (h *handler) Register(c *gin.Context) {
	registerInput := &sto.RegisterUserInput{}
	err := c.ShouldBindJSON(registerInput)
	if err != nil {
		ginresponse.WithError(c, http.StatusUnprocessableEntity, err, "invalid object format")
		return
	}

	id, err := h.authService.RegisterUser(c, registerInput)
	if err != nil {
		ginresponse.WithError(c, http.StatusInternalServerError, err, "error registering user")
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

func (h *handler) Login(c *gin.Context) {
	loginUserInput := sto.LoginUserInput{}
	err := c.ShouldBindJSON(&loginUserInput)
	if err != nil {
		ginresponse.WithError(c, http.StatusUnprocessableEntity, err, "invalid object format")
		return
	}

	jwtToken, refreshToken, err := h.authService.LoginUser(c, &loginUserInput)
	if errors.Is(err, storage.ErrNotFound) {
		ginresponse.WithError(c, http.StatusBadRequest, err, "invalid login or password")
		return
	}

	if err != nil {
		ginresponse.WithError(c, http.StatusInternalServerError, err, "couldn't get user")
		return
	}

	c.JSON(http.StatusOK, gin.H{"Bearer": *jwtToken, "Rt": *refreshToken})
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
