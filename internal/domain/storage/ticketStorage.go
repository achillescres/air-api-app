package storage

import (
	"api-app/internal/domain/dto"
	"api-app/internal/domain/entity"
)

type TicketStorage Storage[entity.Ticket, entity.TicketView, dto.TicketCreate]
