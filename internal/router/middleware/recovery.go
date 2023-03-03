package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/lllllan-fv/gateway-proxy/conf"
	"github.com/lllllan-fv/gateway-proxy/public/resp"
)

// Recovery Capture all panic and return error message
func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				if conf.GetConfig().App.IsProdEnv() {
					// Do not display specific internal errors in the production environment
					resp.DefaultInternalError(c, c.Err())
				} else {
					resp.InternalError(c, c.Err())
				}
			}
		}()
		c.Next()
	}
}
