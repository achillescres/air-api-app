package httpHandler

import (
	"github.com/achillescres/saina-api/internal/config"
	service "github.com/achillescres/saina-api/internal/domain/service"
	httpMiddleware "github.com/achillescres/saina-api/internal/infrastructure/controller/handler/http/middleware"
	"github.com/achillescres/saina-api/internal/infrastructure/gateway/tais"
	"github.com/gin-gonic/gin"
)

type Handler interface {
	RegisterRouter(r *gin.RouterGroup) error
}

type handler struct {
	middleware  httpMiddleware.Middleware
	authService service.AuthService
	taisParser  parser.TaisParser
	taisOutput  parser.TaisOutput
	dataService service.DataService
	cfg         config.HandlerConfig
}

func NewHandler(middleware httpMiddleware.Middleware,
	authService service.AuthService,
	taisParser parser.TaisParser,
	taisOutput parser.TaisOutput,
	dataService service.DataService,
	cfg config.HandlerConfig,
) Handler {
	return &handler{middleware: middleware, authService: authService, taisParser: taisParser, taisOutput: taisOutput, dataService: dataService, cfg: cfg}
}

//func (h *handler) _parse(c *gin.Context) {
//	_, errs, err := h.taisParser.ParseFirstTaisFile(c)
//	if err != nil {
//		ginresponse.JSON(c, http.StatusInternalServerError, gin.H{"error": "error parsing tais file", "errors": errs})
//		return
//	}
//}

func (h *handler) RegisterRouter(r *gin.RouterGroup) error {
	h.registerAuth(r)

	api := r.Group("/api")

	//api.GET("/_parse", h._parse)
	h.registerFlightTable(api)
	h.registerTicket(api)
	h.registerTais(api)

	return nil
}
