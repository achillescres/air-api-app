package product

import (
	"context"
	"github.com/achillescres/saina-api/internal/config"
	"github.com/achillescres/saina-api/internal/infrastructure/controller"
	httpMiddleware "github.com/achillescres/saina-api/internal/infrastructure/controller/handler/http/middleware"
	"github.com/achillescres/saina-api/internal/infrastructure/controller/handler/http/v1"
	"github.com/achillescres/saina-api/pkg/aws/s3"
	"github.com/gin-gonic/gin"
)

type Controllers interface {
	RegisterHandlers(r *gin.RouterGroup) error
	RunTaisController(ctx context.Context) error
}

type controllers struct {
	handler              httpHandler.Handler
	taisParserController controller.TaisParserController
}

// RunTaisController locks current goroutine
func (c *controllers) RunTaisController(ctx context.Context) error {
	err := c.taisParserController.RunFTPWatcher(ctx)
	return err
}

func NewControllers(
	gateways *Gateways,
	bucket s3.Bucket,
	handlerCfg config.HandlerConfig,
	middlewareCfg config.MiddlewareConfig,
	taisParserControllerCfg config.TaisParserControllerConfig,
) (Controllers, error) {
	taisParserController, err := controller.NewTaisParserController(taisParserControllerCfg, gateways.TaisParser, bucket)
	if err != nil {
		return nil, err
	}
	return &controllers{
		handler: httpHandler.NewHandler(
			httpMiddleware.NewMiddleware(middlewareCfg, gateways.Services.AuthService),
			gateways.Services.AuthService,
			gateways.TaisParser,
			gateways.TaisOutput,
			gateways.Services.TablesService,
			handlerCfg,
		),
		taisParserController: taisParserController,
	}, nil
}

func (c *controllers) RegisterHandlers(r *gin.RouterGroup) error {
	return c.handler.RegisterRouter(r)
}
