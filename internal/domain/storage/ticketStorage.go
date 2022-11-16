package storage

import (
	"api-app/internal/domain/entity"
	"api-app/internal/domain/storage/dto"
)

type TicketStorage Storage[entity.Ticket, entity.TicketView, dto.TicketCreate]
