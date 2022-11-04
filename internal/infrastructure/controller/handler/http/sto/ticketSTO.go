package sto

import "api-app/internal/domain/entity"

type TicketSTO entity.Ticket

func ToTicketSTO(t entity.Ticket) TicketSTO {
	return TicketSTO(t)
}

type TicketViewSTO entity.TicketView

func ToTicketViewSTO(t entity.TicketView) TicketViewSTO {
	return TicketViewSTO(t)
}
