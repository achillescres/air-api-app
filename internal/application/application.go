package application

import (
	"api-app/internal/config"
	parser "api-app/internal/infrastructure/controller/parser/filesystem"
	"api-app/internal/product"
	"api-app/pkg/db/postgresql"
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
	"net/http"
)

type App interface {
	Run(ctx context.Context) error
	runHTTP(ctx context.Context) error
}

type app struct {
	cfg        config.AppConfig
	httpServer *product.Routers
	pgPool     postgresql.PGXPool
}

func NewApp(ctx context.Context) (App, error) {
	// Get all needed configs
	log.Infoln("Gathering configs...")
	appCfg := config.App()
	usecaseCfg := config.Usecase()
	parserCfg := config.TaisParser()
	dbCfg := config.Postgres()
	log.Infoln("Success")

	pgPool, err := postgresql.NewPGXPool(ctx, &postgresql.ClientConfig{
		MaxConnections:        dbCfg.MaxConnections,
		MaxConnectionAttempts: dbCfg.MaxConnectionAttempts,
		WaitingDuration:       dbCfg.WaitTimeout,
		Username:              dbCfg.Username,
		Password:              dbCfg.Password,
		Host:                  dbCfg.Host,
		Port:                  dbCfg.Port,
		Database:              dbCfg.Database,
	})
	if err != nil {
		return nil, err
	}

	// Create repositories passing to them database config

	log.Infoln("Creating repository...")
	repos, err := product.NewRepositories(pgPool, &dbCfg)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("fatal couldn't create repositories: %s", err.Error()))
	}
	log.Infoln("Success!")

	log.Infoln("Creating services...")
	services, err := product.NewServices(repos)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("fatal couldn't create services: %s", err.Error()))
	}
	log.Infoln("Success!")

	log.Infoln("Creating usecases...")
	usecases, err := product.NewUsecases(services, &usecaseCfg)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("fatal couldn't create usecases: %s", err.Error()))
	}
	log.Infoln("Success!")

	log.Infoln("Creating handlers...")
	handlers, err := product.NewHandlers(usecases)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("fatal couldn't create handlers: %s", err.Error()))
	}
	log.Infoln("Success!")

	log.Infoln("Building routers...")
	router, err := product.NewRouters(ctx, handlers)
	if err != nil {
		return nil, err
	}

	log.Infoln("Building parsers...")
	taisParser := parser.NewTaisParser(usecases.TicketUc(), usecases.FlightUc(), parserCfg)
	log.Infoln("Success")

	// In this case I distinguished parsers as another type of controllers
	// For now there's only the TAIS File parser, this is executing artificially at start of the app
	// TODO normally IMPLEMENT TAIS FILE PARSING
	//go func() {
	//	err := taisParser.ParseFirstTaisFile(ctx)
	//	if err != nil {
	//		log.Fatalf("fatal INITIAL PARSE tais file: %s\n", err.Error())
	//	}
	//}()
	// ---

	// TODO think what to do with this
	router.GET("/api/_parse", func(c *gin.Context) {
		err := taisParser.ParseFirstTaisFile(c)
		if err != nil {
			log.Errorf("error parsing tais file: %s\n", err.Error())
			c.AbortWithStatus(http.StatusInternalServerError)
		}
	})

	return &app{
		httpServer: router,
		cfg:        appCfg,
		pgPool:     pgPool,
	}, nil
}

func (app *app) Run(ctx context.Context) error {
	grp, ctx := errgroup.WithContext(ctx)

	grp.Go(func() error {
		return app.runHTTP(ctx)
	})

	return grp.Wait()
}

func (app *app) runHTTP(ctx context.Context) error {
	listen := app.cfg.HTTP

	addr := listen.IP
	if len(listen.Port) != 0 {
		addr += ":" + listen.Port
	}

	err := app.httpServer.Run(addr)
	if err != nil {
		log.Errorf("error unable ot run httpServer\n")
	} else {
		log.Infof("Listening to %s\n", addr)
	}
	return err
}
