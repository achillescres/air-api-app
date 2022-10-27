package main

import (
	"api-app/internal/handler"
	parser "api-app/internal/parser/filesystem"
	"api-app/internal/usecase"
	"api-app/internal/usecase/composite"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

const TAIS_PATH = "C:\\TinyGainAir\\api-app\\external\\tais.txt"

func main() {
	const ADDR = "127.0.0.1:777"

	log.Infoln("Starting...")
	r := gin.Default()

	log.Infoln("Building inner layers...")
	flightService := composite.GenerateFlightComposite()
	tiketService := composite.GenerateTicketComposite()

	flightUc := usecase.NewFlightUsecase(flightService, tiketService)
	ticketUc := usecase.NewTicketUsecase(tiketService)

	flightHandler := handler.NewFlightHandler(flightUc)
	//ticketHandler := handler.NewTicketHandler(ticketUc)

	taisParser := parser.NewTaisParser(ticketUc, flightUc)
	err := taisParser.ParseFile(TAIS_PATH)

	if err != nil {
		log.Errorf("error parsing tais file: %s", err.Error())
		return
	}

	log.Infoln("Attaching routers...")
	r.GET("/api/getAllFlightTables", flightHandler.GetAllFlightTables)
	r.POST("/api/_parse", func(c *gin.Context) {
		err := taisParser.ParseFile(TAIS_PATH)
		if err != nil {
			log.Errorf("error parsing tais file: %s", err.Error())
			c.JSON(http.StatusInternalServerError, nil)
			return
		}
	})

	err = r.Run(ADDR)
	if err != nil {
		log.Errorf("error can't run server: %s", err.Error())
		os.Exit(1)
		return
	}
}
