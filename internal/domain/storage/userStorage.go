package storage

import (
	"api-app/internal/domain/dto"
	"api-app/internal/domain/entity"
)

type UserStorage Storage[entity.User, entity.UserView, dto.UserCreate]
