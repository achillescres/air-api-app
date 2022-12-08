package httpHandler

import (
	"github.com/achillescres/saina-api/internal/config"
	service2 "github.com/achillescres/saina-api/internal/domain/service"
	httpMiddleware "github.com/achillescres/saina-api/internal/infrastructure/controller/handler/http/middleware"
	"github.com/achillescres/saina-api/internal/infrastructure/controller/parser/filesystem"
	"github.com/achillescres/saina-api/pkg/gin/ginresponse"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler interface {
	RegisterRouter(r *gin.RouterGroup) error
}

type handler struct {
	middleware  httpMiddleware.Middleware
	authService service2.AuthService
	taisParser  parser.TaisParser
	dataService service2.DataService
	cfg         config.HandlerConfig
}

func NewHandler(middleware httpMiddleware.Middleware, authService service2.AuthService, taisParser parser.TaisParser, dataService service2.DataService, cfg config.HandlerConfig) *handler {
	return &handler{middleware: middleware, authService: authService, taisParser: taisParser, dataService: dataService, cfg: cfg}
}

var _ Handler = (*handler)(nil)

func (h *handler) _parse(c *gin.Context) {
	_, _, err := h.taisParser.ParseFirstTaisFile(c)
	if err != nil {
		ginresponse.WithError(c, http.StatusInternalServerError, err, "couldn't parse tais file")
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
