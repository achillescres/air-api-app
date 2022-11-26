package gintools

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func ErrorToContext(c *gin.Context, err error) {
	err = c.Error(err)
	if err != nil {
		log.Errorf("error puting error to gin context: %s\n", err.Error())
	}
}
