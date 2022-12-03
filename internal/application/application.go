package application

import (
	"context"
	"errors"
	"fmt"
	config2 "github.com/achillescres/saina-api/internal/config"
	"github.com/achillescres/saina-api/internal/infrastructure/controller/parser/filesystem"
	product2 "github.com/achillescres/saina-api/internal/product"
	postgresql2 "github.com/achillescres/saina-api/pkg/db/postgresql"
	"github.com/achillescres/saina-api/pkg/security/ajwt"
	"github.com/achillescres/saina-api/pkg/security/passlib"
	log "github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
)

type App interface {
	Run(ctx context.Context) error
	runHTTP() error
}

type app struct {
	cfg        config2.AppConfig
	httpServer *product2.Routers
	pgPool     postgresql2.PGXPool
}

func NewApp(ctx context.Context) (App, error) {
	// Get all needed configs
	log.Infoln("Gathering configs...")
	envCfg := config2.Env()
	appCfg := config2.App()
	handlerCfg := config2.Handler()
	taisParserCfg := config2.TaisParser()
	dbCfg := config2.Postgres()
	authCfg := config2.Auth()
	log.Infoln("Success!")

	log.Warnln(envCfg.ProjectAbsPath)
	hashManager := passlib.NewHashManager(envCfg.PasswordHashSalt)
	jwtManager := ajwt.NewJWTManager(hashManager, envCfg.JWTSecret, authCfg.JWTLiveTime, authCfg.RefreshTokenLiveTime)
	pgPool, err := postgresql2.NewPGXPool(ctx, &postgresql2.ClientConfig{
		MaxConnections:        dbCfg.MaxConnections,
		MaxConnectionAttempts: dbCfg.MaxConnectionAttempts,
		WaitingDuration:       dbCfg.WaitTimeout,
		Username:              dbCfg.Username,
		Password:              dbCfg.Password,
		Host:                  dbCfg.Host,
		Port:                  dbCfg.Port,
		Database:              dbCfg.Database,
	})
	log.Warnln(pgPool)
	if err != nil {
		return nil, err
	}

	// Create repositories passing to them database config
	log.Infoln("Creating repository...")
	repos, err := product2.NewRepositories(pgPool, hashManager)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("fatal couldn't create repositories: %s", err.Error()))
	}
	log.Infoln("Success!")

	log.Infoln("Creating services...")
	services, err := product2.NewServices(repos, &taisParserCfg, hashManager, jwtManager, authCfg)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("fatal couldn't create services: %s", err.Error()))
	}
	log.Infoln("Success!")

	log.Infoln("Creating filesystem controllers")
	taisParser := parser.NewTaisParser(services.ParserService, taisParserCfg)
	log.Infoln("Success!")

	log.Infoln("Creating internet controllers(handlers)...")
	handlers, err := product2.NewHandlers(services, &handlerCfg, taisParser)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("fatal couldn't create handlers: %s", err.Error()))
	}
	log.Infoln("Success!")

	log.Infoln("Building routers...")
	router, err := product2.NewRouters(handlers)
	if err != nil {
		return nil, err
	}

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
	//router.GET("/api/_parse", func(c *gin.Context) {
	//	err := taisParser.ParseFirstTaisFile(c)
	//	if err != nil {
	//		log.Errorf("error parsing tais file: %s\n", err.Error())
	//		c.AbortWithStatus(http.StatusInternalServerError)
	//	}
	//})

	return &app{
		httpServer: router,
		cfg:        appCfg,
		pgPool:     pgPool,
	}, nil
}

func (app *app) Run(ctx context.Context) error {
	grp, ctx := errgroup.WithContext(ctx)

	grp.Go(func() error {
		return app.runHTTP()
	})

	return grp.Wait()
}

func (app *app) runHTTP() error {
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
