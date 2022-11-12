package storage

import "api-app/internal/domain/entity"

type TicketStorage Storage[entity.Ticket, entity.TicketView]
