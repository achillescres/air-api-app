package service

import (
	"api-app/internal/domain/dto"
	"api-app/internal/domain/storage"
	"api-app/pkg/gin/ginresponse"
	"context"
	"net/http"
)

type AuthService interface {
	RegisterUser(regInput dto.RegisterInput) error
}

type authService struct {
	userStorage storage.UserStorage
}

func (aS *authService) RegisterUser(ctx context.Context, regInput dto.RegisterInput) error {
	user, err := aS.userStorage.Store(c, createUser)
	if err != nil {
		ginresponse.WithError(c, http.StatusInternalServerError, err, "couldn't store user")
	}
}
