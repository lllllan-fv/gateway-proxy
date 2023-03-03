package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lllllan-fv/gateway-proxy/internal/router/middleware"
	"github.com/lllllan-fv/gateway-proxy/public/resp"
)

func InitHttpRouter(middlewares ...gin.HandlerFunc) *gin.Engine {
	router := gin.New()
	router.Use(middlewares...)

	router.GET("/ping", func(c *gin.Context) { resp.Success(c, "pong") })

	router.Use(
		middleware.HTTPAccessModeMiddleware(),
	)

	return router
}
