package storage

import (
	"api-app/internal/domain/entity"
	"api-app/internal/domain/storage/dto"
)

type UserStorage Storage[entity.User, entity.UserView, dto.UserCreate]
