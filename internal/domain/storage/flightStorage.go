package storage

import (
	"api-app/internal/domain/dto"
	"api-app/internal/domain/entity"
)

type FlightStorage Storage[entity.Flight, entity.FlightView, dto.FLightCreate]
