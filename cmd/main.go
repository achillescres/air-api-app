package main

import (
	"api-app/internal/infrastructure/handler"
	parser "api-app/internal/infrastructure/parser/fsParser"
	"api-app/internal/usecase"
	"api-app/internal/usecase/composite"
	"api-app/pkg/logging"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

const TAISPATH = "C:\\TinyGainAir\\api-app\\external\\tais.txt"

func init() {
	logging.ConfigureLog()
}

func main() {
	log.Infoln("Starting...")
	//conf := config.GetMainConfig()
	//fmt.Println(conf)

	log.Infoln("Building inner layers...")
	flightService := composite.GenerateFlightComposite()
	ticketService := composite.GenerateTicketComposite()

	flightUc := usecase.NewFlightUsecase(flightService, ticketService)
	ticketUc := usecase.NewTicketUsecase(ticketService)

	flightHandler := handler.NewFlightHandler(flightUc)
	//ticketHandler := handler.NewTicketHandler(ticketUc)

	taisParser := parser.NewTaisParser(ticketUc, flightUc)
	err := taisParser.ParseFile(TAISPATH)

	if err != nil {
		log.Errorf("error parsing tais file: %s", err.Error())
		return
	}

	log.Infoln("Attaching routers...")
	r := gin.Default()

	r.GET("/api/getAllFlightTables", flightHandler.GetAllFlightTables)

	r.POST("/api/_parse", func(c *gin.Context) {
		err := taisParser.ParseFile(TAISPATH)
		if err != nil {
			log.Errorf("error parsing tais file: %s", err.Error())
			c.JSON(http.StatusInternalServerError, nil)
			return
		}
	})

	err = r.Run(
		"127.0.0.1:7771",
	)
	if err != nil {
		log.Fatalf("error can't run server: %s", err.Error())
		return
	}
}
