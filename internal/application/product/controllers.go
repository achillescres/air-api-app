package product

import (
	"github.com/achillescres/saina-api/internal/config"
	"github.com/achillescres/saina-api/internal/infrastructure/controller/handler/http/v1"
	parser "github.com/achillescres/saina-api/internal/infrastructure/controller/parser/tais"
	"github.com/gin-gonic/gin"
)

type Controllers interface {
	Register(r *gin.RouterGroup) error
}

type controllers struct {
	handler              httpHandler.Handler
	taisParserController parser.TaisParserController
}

func NewControllers(
	services *Services,
	middlewares *Middlewares,
	taisParserCfg config.TaisParserConfig,
	taisParserControllerCfg config.TaisParserControllerConfig,
	handlerCfg *config.HandlerConfig,
) (Controllers, error) {
	taisParser := parser.NewTaisParser(services.ParserService, taisParserCfg)
	taisParserController := parser.NewTaisParserController(taisParserControllerCfg, taisParser)
	return &controllers{
		handler: httpHandler.NewHandler(
			middlewares.middleware,
			services.AuthService,
			taisParser,
			services.TablesService,
			*handlerCfg,
		),
		taisParserController: taisParserController,
	}, nil
}

func (h *controllers) Register(r *gin.RouterGroup) error {
	return h.handler.RegisterRouter(r)
}
