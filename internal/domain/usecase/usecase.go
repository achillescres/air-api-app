package usecase

import (
	"api-app/internal/domain/entity"
	"api-app/internal/domain/service"
	"api-app/internal/domain/storage/dto"
)

type Usecase[Entity entity.Entity, View entity.View, Create dto.Create] interface {
	service.Service[Entity, View, Create]
}
