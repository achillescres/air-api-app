package httpMiddleware

import (
	"errors"
	"github.com/achillescres/saina-api/pkg/gin/ginresponse"
	"github.com/achillescres/saina-api/pkg/object/oid"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

var (
	errEmptyHeader = errors.New(
		"user error auth header is empty",
	)
	errInvalidHeader = errors.New(
		"user error invalid auth header",
	)
)

func (m *middleware) ParseAndInjectTokenMiddleware(c *gin.Context) {
	header := c.GetHeader(m.middlewareConfig.AuthorizationHeader)
	if header == "" {
		ginresponse.WithError(c, http.StatusUnauthorized, errEmptyHeader, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		ginresponse.WithError(c, http.StatusUnauthorized, errInvalidHeader, "invalid auth header")
	}

	if headerParts[1] == "" {
		ginresponse.WithError(c, http.StatusUnauthorized, errInvalidHeader, "token is empty")
	}
	token := headerParts[1]

	userId, err := m.authService.ParseUserToken(c, token)
	if err != nil {
		ginresponse.WithError(c, http.StatusUnauthorized, err, "couldn't parse token")
	}

	c.Set(m.middlewareConfig.UserIdCtxKey, userId)
}

func (m *middleware) GetUserId(c *gin.Context) (oid.Id, error) {
	id, ok := c.Get(m.middlewareConfig.UserIdCtxKey)
	if !ok {
		return oid.Undefined, errors.New("userId not found")
	}

	nId, ok := oid.AssertId(id)
	if !ok {
		return oid.Undefined, errors.New("userId is of invalid type")
	}

	return nId, nil
}
