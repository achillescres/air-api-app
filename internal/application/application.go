package application

import (
	"context"
	"fmt"
	"github.com/achillescres/saina-api/internal/application/product"
	"github.com/achillescres/saina-api/internal/config"
	"github.com/achillescres/saina-api/pkg/aws/s3"
	"github.com/achillescres/saina-api/pkg/db/postgresql"
	"github.com/achillescres/saina-api/pkg/security/ajwt"
	"github.com/achillescres/saina-api/pkg/security/passlib"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
)

type App interface {
	Run(ctx context.Context) error
	runHTTP() error
	runTaisParsing(ctx context.Context) error
}

type app struct {
	cfg      config.AppConfig
	pgPool   postgresql.PGXPool
	endpoint product.Endpoint
}

func NewApp(ctx context.Context) (App, error) {
	// Get env config
	log.Infoln("Gathering all configs")

	// Outer managers
	log.Infoln("Creating managers:")
	hashManager := passlib.NewHashManager(config.Env().PasswordHashSalt)
	jwtManager := ajwt.NewJWTManager(hashManager, config.Env().JWTSecret, config.Auth().JWTLiveTime, config.Auth().RefreshTokenLiveTime)

	// Build postgres pool
	log.Infoln("Creating pgxpool")
	pgPool, err := postgresql.NewPGXPool(ctx, &postgresql.ClientConfig{
		MaxConnections:        config.Postgres().MaxConnections,
		MaxConnectionAttempts: config.Postgres().MaxConnectionAttempts,
		WaitingDuration:       config.Postgres().WaitTimeout,
		Username:              config.Postgres().Username,
		Password:              config.Postgres().Password,
		Host:                  config.Postgres().Host,
		Port:                  config.Postgres().Port,
		Database:              config.Postgres().Database,
	})
	if err != nil {
		log.Errorf("error creating pgxpool: %s\n", err)
		//return nil, err
	}
	log.Warnf("This is pgxpool: %s", pgPool)

	log.Infoln("Creating bucket connection")
	log.Warnln(config.TaisParserController().BucketName)
	bucket, err := s3.NewBucket(ctx,
		config.TaisParserController().BucketName,
		config.TaisParserController().FileDownloadTL,
		config.TaisParserController().FileUploadTL,
	)
	if err != nil {
		return nil, err
	}

	// Create repositories
	log.Infoln("Creating repository...")
	repos, err := product.NewRepositories(pgPool, hashManager)
	if err != nil {
		return nil, fmt.Errorf("fatal couldn't create repositories: %s", err)
	}
	log.Infoln("Success!")

	log.Infoln("Creating services...")
	services, err := product.NewServices(repos, config.Tais(), hashManager, jwtManager, config.Auth())
	if err != nil {
		return nil, fmt.Errorf("fatal couldn't create services: %s", err)
	}
	log.Infoln("Success!")

	// Outer layer
	log.Infoln("Controllers layer:")

	// Worker sub-layer
	log.Infoln("Gateways sub-layer")
	gateways := product.NewGateways(services, config.Tais(), bucket)
	log.Infoln("Success!")

	// Controllers sub-layer
	log.Infoln("Creating controllers")
	controllers, err := product.NewControllers(gateways, bucket, config.Handler(), config.Middleware(), config.TaisParserController())
	if err != nil {
		return nil, fmt.Errorf("fatal couldn't create controllers: %s", err)
	}
	log.Infoln("Success!")

	// Endpoint artificial-layer
	log.Infoln("Creating artificial endpoint layer")
	endpoint := product.NewEndpoint(controllers)
	log.Infoln("Success")

	return &app{
		cfg:      config.App(),
		endpoint: endpoint,
		pgPool:   pgPool,
	}, nil
}

func (app *app) Run(ctx context.Context) error {
	grp, ctx := errgroup.WithContext(ctx)

	grp.Go(func() error {
		return app.runHTTP()
	})

	grp.Go(func() error {
		return app.runTaisParsing(ctx)
	})

	return grp.Wait()
}

func (app *app) runHTTP() error {
	r := gin.Default()
	err := app.endpoint.RegisterHandlersToGroup(r)
	if err != nil {
		return err
	}

	addr := app.cfg.HTTP.IP
	if len(app.cfg.HTTP.Port) != 0 {
		addr += ":" + app.cfg.HTTP.Port
	}

	err = r.Run(addr)
	if err != nil {
		log.Errorf("error unable ot run httpServer\n")
	}
	return err
}

func (app *app) runTaisParsing(ctx context.Context) error {
	defer func() {
		if r := recover(); r != nil {
			err := app.runTaisParsing(ctx)
			if err != nil {
				err = fmt.Errorf("error couldn't recover TaisParsing after panic: %s", err)
				log.Errorln(err)
				panic(err)
			}
		}
	}()
	err := app.endpoint.RunTaisParserController(ctx)
	return err
}
