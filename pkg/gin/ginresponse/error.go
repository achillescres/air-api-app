package ginresponse

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type errorJSON struct {
	Error string `json:"error"`
}

func Error(c *gin.Context, code int, err error, resError string) {
	log.Errorln(err)

	c.AbortWithStatusJSON(code, errorJSON{Error: resError})
	if err := c.Error(err); err != nil {
		log.Errorln(err)
	}
}

func JSON(c *gin.Context, code int, json gin.H) {
	c.JSON(code, json)
	if err := c.Err(); err != nil {
		log.Errorf("error occured in gin context: %s", err)
	}
}
