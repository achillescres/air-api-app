package usecase

import (
	"api-app/internal/domain/entity"
	"api-app/internal/domain/service"
)

type Usecase[Entity entity.Entity, View entity.View] interface {
	service.Service[Entity, View]
}
