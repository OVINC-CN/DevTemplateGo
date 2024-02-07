package middlewares

import (
	"github.com/OVINC-CN/DevTemplateGo/src/configs"
	ginI18n "github.com/gin-contrib/i18n"
	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Timeout() gin.HandlerFunc {
	return timeout.New(
		timeout.WithTimeout(configs.Config.RequestTimeout),
		timeout.WithHandler(func(c *gin.Context) {
			c.Next()
		}),
		timeout.WithResponse(func(c *gin.Context) {
			c.AbortWithStatusJSON(
				http.StatusGatewayTimeout,
				gin.H{
					"error": gin.H{
						"message": ginI18n.MustGetMessage(c, "timeout"),
					},
				},
			)
		}),
	)
}
