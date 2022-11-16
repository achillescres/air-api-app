package storage

import (
	"api-app/internal/domain/entity"
	"api-app/internal/domain/storage/dto"
)

type FlightStorage Storage[entity.Flight, entity.FlightView, dto.FLightCreate]
