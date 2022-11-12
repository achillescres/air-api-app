package main

import (
	"api-app/internal/config"
	"api-app/internal/domain/service"
	"api-app/internal/domain/usecase"
	"api-app/internal/infrastructure/controller/handler/http"
	"api-app/internal/infrastructure/controller/parser/filesystem"
	"api-app/internal/infrastructure/controller/router"
	"api-app/internal/infrastructure/repository"
	"api-app/pkg/logging"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func init() {
	logging.ConfigureLog()
}

func main() {
	log.Infoln("Gathering configs...")
	appCfg := config.App()
	usecaseCfg := config.Usecase()
	parserCfg := config.TaisParser()
	log.Infoln("Success")

	log.Infoln("Building repositories...")
	flightRepo := repository.NewFlightRepository()
	ticketRepo := repository.NewTicketRepository()
	userRepo := repository.NewUserRepository()
	log.Infoln("Success")

	log.Infoln("Building services...")
	flightService := service.NewFlightService(flightRepo)
	ticketService := service.NewTicketService(ticketRepo)
	userService := service.NewUserService(userRepo)
	log.Infoln("Success")

	log.Infoln("Building usecases...")
	flightUc := usecase.NewFlightUsecase(flightService, ticketService, usecaseCfg)
	ticketUc := usecase.NewTicketUsecase(ticketService)
	userUc := usecase.NewUserUsecase(userService)
	log.Infoln("Success")

	log.Infoln("Building controllers:")
	log.Infoln("Building handlers...")
	flightHandler := httpHandler.NewFlightHandler(flightUc)
	ticketHandler := httpHandler.NewTicketHandler(ticketUc)
	userHandler := httpHandler.NewUserHandler(userUc)
	log.Infoln("Success")

	log.Infoln("Building parsers...")
	taisParser := parser.NewTaisParser(ticketUc, flightUc, parserCfg)
	log.Infoln("Success")

	// TODO implement normal initial parse
	err := taisParser.ParseFirstTaisFile()
	if err != nil {
		log.Errorf("error parsing tais file: %s\n", err.Error())
		return
	}
	// ---

	log.Infoln("Building routers...")
	r := gin.Default()

	router.RegisterFlightRouter(r, flightHandler)
	router.RegisterTicketRouter(r, ticketHandler)
	router.RegisterUserRouter(r, userHandler)

	r.POST("/api/_parse", func(c *gin.Context) { // TODO think what to do with this
		err := taisParser.ParseFirstTaisFile()
		if err != nil {
			log.Errorf("error parsing tais file: %s\n", err.Error())
			c.JSON(http.StatusInternalServerError, nil)
			return
		}
	})

	listen := appCfg.HTTP
	err = r.Run(
		fmt.Sprintf("%s:%s", listen.IP, listen.Port),
	)
	if err != nil {
		log.Fatalf("error can't run server: %s\n", err.Error())
	}
}
