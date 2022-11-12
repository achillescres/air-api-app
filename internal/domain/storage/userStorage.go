package storage

import "api-app/internal/domain/entity"

type UserStorage Storage[entity.User, entity.UserView]
