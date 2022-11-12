package router

import (
	httpHandler "api-app/internal/infrastructure/controller/handler/http"
	"api-app/pkg/rfmt"
	"github.com/gin-gonic/gin"
)

func RegisterUserRouter(r *gin.Engine, handler httpHandler.UserHandler) *gin.Engine {
	r.GET(
		rfmt.JoinApiRoute("login"),
		handler.Login,
	)

	r.GET(
		rfmt.JoinApiRoute("logout"),
		handler.Logout,
	)

	r.POST(
		rfmt.JoinApiRoute("register"),
		handler.Register,
	)

	return r
}
