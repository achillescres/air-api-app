package main

import (
	"api-app/internal/config"
	"api-app/internal/domain/service"
	"api-app/internal/domain/usecase"
	"api-app/internal/infrastructure/controller/handler/http"
	"api-app/internal/infrastructure/controller/parser/filesystem"
	"api-app/internal/infrastructure/repository"
	"api-app/pkg/logging"
	"fmt"
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
	log.Infoln("Getting configs")
	_, appCfg, usecaseCfg, parserCfg, err := config.OnceGetConfigInvokes()
	if err != nil {
		log.Fatalf("fatal OnceGetConfigInvokes: %s\n", err.Error())
	}

	log.Infoln("Building inner layers...")
	flightRepo := repository.NewFlightRepository()
	ticketRepo := repository.NewTicketRepository()
	flightService := service.NewFlightService(flightRepo)
	ticketService := service.NewTicketService(ticketRepo)

	flightUc := usecase.NewFlightUsecase(flightService, ticketService, usecaseCfg)
	ticketUc := usecase.NewTicketUsecase(ticketService)

	flightHandler := httpHandler.NewFlightHandler(flightUc)
	ticketHandler := httpHandler.NewTicketHandler(ticketUc)

	taisParser := parser.NewTaisParser(ticketUc, flightUc, parserCfg)
	err = taisParser.ParseTaisFile()

	if err != nil {
		log.Errorf("error parsing tais file: %s\n", err.Error())
		return
	}

	log.Infoln("Attaching routers...")
	r := gin.Default()

	r.GET("/api/getAllFlightTables", flightHandler.GetAllFlightTables)

	r.POST("/api/_parse", func(c *gin.Context) {
		err := taisParser.ParseTaisFile()
		if err != nil {
			log.Errorf("error parsing tais file: %s\n", err.Error())
			c.JSON(http.StatusInternalServerError, nil)
			return
		}
	})

	listen := appCfg().Listen
	err = r.Run(
		fmt.Sprintf("%s:%s", listen.IP, listen.Port),
	)
	if err != nil {
		log.Fatalf("error can't run server: %s\n", err.Error())
		return
	}
}
