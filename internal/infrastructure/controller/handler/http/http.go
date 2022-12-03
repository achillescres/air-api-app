package httpHandler

import (
	"github.com/achillescres/saina-api/internal/config"
	service2 "github.com/achillescres/saina-api/internal/domain/service"
	"github.com/achillescres/saina-api/internal/infrastructure/controller/parser/filesystem"
	"github.com/achillescres/saina-api/pkg/gin/ginresponse"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler interface {
	RegisterRouter(r *gin.RouterGroup) error
}

type handler struct {
	authService service2.AuthService
	taisParser  parser.TaisParser
	dataService service2.DataService
	cfg         config.HandlerConfig
}

var _ Handler = (*handler)(nil)

func NewHandler(
	authService service2.AuthService,
	parserService parser.TaisParser,
	dataService service2.DataService,
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
