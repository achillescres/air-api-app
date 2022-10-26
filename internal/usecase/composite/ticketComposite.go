package composite

import (
	"api-app/internal/adapter/service"
	"api-app/internal/repository"
)

func GenerateTicketComposite() service.TicketService {
	repo := repository.NewTicketRepository()
	serv := service.NewTicketService(repo)
	return serv
}
