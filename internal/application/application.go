package application

import (
	"context"
	"errors"
	"fmt"
	"github.com/achillescres/saina-api/internal/application/product"
	"github.com/achillescres/saina-api/internal/config"
	"github.com/achillescres/saina-api/internal/infrastructure/controller/parser/filesystem"
	"github.com/achillescres/saina-api/pkg/db/postgresql"
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
	cfg        config.AppConfig
	httpServer *product.Routers
	pgPool     postgresql.PGXPool
}

func NewApp(ctx context.Context) (App, error) {
	// Get env config
	log.Infoln("Gathering all configs")
	envCfg := config.Env()
	authCfg := config.Auth()
	dbCfg := config.Postgres()
	taisParserCfg := config.TaisParser()
	middlewareCfg := config.Middleware()
	handlerCfg := config.Handler()
	appCfg := config.App()

	// Outer/Atomic layer managers
	log.Infoln("Creating managers:")
	hashManager := passlib.NewHashManager(envCfg.PasswordHashSalt)
	jwtManager := ajwt.NewJWTManager(hashManager, envCfg.JWTSecret, authCfg.JWTLiveTime, authCfg.RefreshTokenLiveTime)

	// Build postgres pool
	log.Infoln("Creating pgxpool")
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
	log.Infof("This is pgxpool: %s", pgPool)

	// Create repositories
	log.Infoln("Creating repository...")
	repos, err := product.NewRepositories(pgPool, hashManager)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("fatal couldn't create repositories: %s", err.Error()))
	}
	log.Infoln("Success!")

	log.Infoln("Creating services...")
	services, err := product.NewServices(repos, &taisParserCfg, hashManager, jwtManager, authCfg)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("fatal couldn't create services: %s", err.Error()))
	}
	log.Infoln("Success!")

	// Controllers layer
	log.Infoln("Layer controllers:")

	// Middleware additional-layer
	log.Infoln("Creating middlewares")
	middleware := product.NewMiddlewares(middlewareCfg, services.AuthService)

	// Parser sub-layer
	log.Infoln("Creating parsers")
	taisParser := parser.NewTaisParser(services.ParserService, taisParserCfg)
	log.Infoln("Success!")

	// Handler sub-layer
	log.Infoln("Creating handlers")
	handlers, err := product.NewHandlers(middleware, services, &handlerCfg, taisParser)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("fatal couldn't create handlers: %s", err.Error()))
	}
	log.Infoln("Success!")

	// External framework layer
	// Routers
	log.Infoln("Building routers...")
	router, err := product.NewRouters(handlers)
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
