package httpHandler

import (
	"github.com/achillescres/saina-api/internal/config"
	service "github.com/achillescres/saina-api/internal/domain/service"
	httpMiddleware "github.com/achillescres/saina-api/internal/infrastructure/controller/handler/http/middleware"
	"github.com/achillescres/saina-api/internal/infrastructure/controller/parser/tais"
	"github.com/achillescres/saina-api/pkg/gin/ginresponse"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler interface {
	RegisterRouter(r *gin.RouterGroup) error
}

type handler struct {
	middleware  httpMiddleware.Middleware
	authService service.AuthService
	taisParser  parser.TaisParser
	dataService service.DataService
	cfg         config.HandlerConfig
}

func NewHandler(middleware httpMiddleware.Middleware, authService service.AuthService, taisParser parser.TaisParser, dataService service.DataService, cfg config.HandlerConfig) *handler {
	return &handler{middleware: middleware, authService: authService, taisParser: taisParser, dataService: dataService, cfg: cfg}
}

var _ Handler = (*handler)(nil)

func (h *handler) _parse(c *gin.Context) {
	_, errs, err := h.taisParser.ParseFirstTaisFile(c)
	if err != nil {
		ginresponse.JSON(c, http.StatusInternalServerError, gin.H{"error": "error parsing tais file", "errors": errs})
		return
	}
}

func (h *handler) RegisterRouter(r *gin.RouterGroup) error {
	auth := r.Group("/auth")
	h.registerAuth(auth)

	api := r.Group("/api") //h.middleware.ParseAndInjectTokenMiddleware)
	h.registerFlightTable(api)
	h.registerTicket(api)
	api.GET("/_parse", h._parse)

	return nil
}
