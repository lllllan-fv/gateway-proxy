package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/lllllan-fv/gateway-proxy/internal/service"
	"github.com/lllllan-fv/gateway-proxy/public/resp"
)

// Match access mode, based on request information
func HTTPAccessModeMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		service, err := service.HTTPAccessMode(c)
		if err != nil {
			resp.Error(c, 1001, err)
			return
		}

		c.Set("service", service)
		c.Next()
	}
}
