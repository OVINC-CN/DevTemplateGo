package middlewares

import (
	"github.com/OVINC-CN/DevTemplateGo/src/core"
	ginI18n "github.com/gin-contrib/i18n"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if p := recover(); p != nil {
				switch err := p.(type) {
				case *core.APIError:
					message, translateError := ginI18n.GetMessage(c, err.Error())
					if translateError != nil {
						message = err.Error()
					}
					c.AbortWithStatusJSON(
						err.Status,
						gin.H{
							"error": gin.H{
								"message": message,
								"detail":  err.Detail,
							},
						},
					)
				case error:
					message, translateError := ginI18n.GetMessage(c, err.Error())
					if translateError != nil {
						message = err.Error()
					}
					c.AbortWithStatusJSON(
						http.StatusInternalServerError,
						gin.H{
							"error": gin.H{
								"message": message,
								"detail":  &map[string]any{},
							},
						},
					)
				default:
					c.AbortWithStatus(http.StatusInternalServerError)
				}
			}
		}()
		c.Next()
	}
}
