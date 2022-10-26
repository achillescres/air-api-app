package main

import (
	"api-app/internal/handler"
	parser "api-app/internal/parser/filesystem"
	"api-app/internal/usecase"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {
	const ADDR = "127.0.0.1:777"

	log.Infoln("Starting...")
	r := gin.Default()

	log.Infoln("Building inner layers...")
	flightUc := usecase.GenerateFlightUsecase()
	ticketUc := usecase.GenerateTicketUsecase()

	flightHandler := handler.NewFlightHandler(flightUc)
	ticketHandler := handler.NewTicketHandler(ticketUc)

	taisParser := parser.NewTaisParser(ticketUc, flightUc)
	err := taisParser.ParseFile("../external/tais.txt")
	if err != nil {
		log.Errorf("error parsing tais file: %s", err.Error())
		return
	}

	log.Infoln("Attaching routers...")
	r.GET("/api/getAllFlightTables", flightHandler.GetAllFlightTables)

	err = r.Run(ADDR)
	if err != nil {
		log.Errorf("error can't run server: %s", err.Error())
		os.Exit(1)
		return
	}
}
