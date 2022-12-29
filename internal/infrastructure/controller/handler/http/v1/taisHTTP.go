package httpHandler

import (
	"fmt"
	"github.com/achillescres/saina-api/internal/infrastructure/controller/safeObject"
	"github.com/achillescres/saina-api/pkg/gin/ginresponse"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func generateTaisOutFilename() string {
	y, m, d := time.Now().Date()
	h, min, s := time.Now().Clock()
	return fmt.Sprintf("OUT_TAIS_FILE_%d-%d-%d_%d:%d:%d", y, m, d, h, min, s)
}

func (h *handler) TaisChanges(c *gin.Context) {
	changes := &safeObject.TaisChangesSafe{}
	err := c.ShouldBindJSON(changes)
	if err != nil {
		ginresponse.Error(c, http.StatusUnprocessableEntity, err, "invalid object format")
		return
	}

	name := generateTaisOutFilename()
	err = h.taisOutput.SendOutputTais(c, name, changes.ToEntity())
	if err != nil {
		ginresponse.Error(c, http.StatusInternalServerError, err, "error sending output file to trade server")
		return
	}
}

func (h *handler) registerTais(r *gin.RouterGroup) {
	r = r.Group("/tais")
	r.POST("/taisChanges", h.TaisChanges)
}
