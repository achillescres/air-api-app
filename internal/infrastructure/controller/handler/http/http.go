package httpHandler

import (
	"api-app/internal/config"
	"api-app/internal/domain/service"
	parser "api-app/internal/infrastructure/controller/parser/filesystem"
	"api-app/pkg/gin/ginresponse"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler interface {
	RegisterRouter(r *gin.RouterGroup) error
}

type handler struct {
	authService service.AuthService
	taisParser  parser.TaisParser
	dataService service.DataService
	cfg         config.HandlerConfig
}

var _ Handler = (*handler)(nil)

func NewHandler(
	authService service.AuthService,
	parserService parser.TaisParser,
	dataService service.DataService,
	cfg config.HandlerConfig,
) Handler {
	return &handler{
		authService: authService,
		taisParser:  parserService,
		dataService: dataService,
		cfg:         cfg}
}

func (h *handler) _parse(c *gin.Context) {
	_, err := h.taisParser.ParseFirstTaisFile(c)
	if err != nil {
		ginresponse.WithError(c, http.StatusInternalServerError, err, "couldn't parse tais file")
		return
	}
}

func (h *handler) RegisterRouter(r *gin.RouterGroup) error {
	auth := r.Group("/auth")
	h.registerUser(auth)

	api := r.Group("/api")
	h.registerFlightTable(api)
	h.registerTicket(api)
	api.GET("/_parse", h._parse)

	return nil
}
