package storage

import (
	"github.com/achillescres/saina-api/internal/domain/entity"
	"github.com/achillescres/saina-api/internal/domain/storage/dto"
)

type FlightStorage Storage[entity.Flight, dto.FLightCreate]
