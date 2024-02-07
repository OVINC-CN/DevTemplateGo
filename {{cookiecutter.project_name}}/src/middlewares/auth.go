package middlewares

import (
	"github.com/OVINC-CN/DevTemplateGo/src/configs"
	"github.com/OVINC-CN/DevTemplateGo/src/core"
	"github.com/OVINC-CN/DevTemplateGo/src/services/account"
	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取用户身份
		sessionID, err := c.Cookie(configs.Config.SessionCookieName)
		if err == nil {
			user := account.LoadUserBySessionID(sessionID)
			if user.Enabled {
				c.Set("User", &user)
			}
		}
		c.Next()
	}
}

func LoginRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := account.GetContextUser(c)
		if user.Username == "" {
			panic(core.LoginRequired)
		}
		c.Next()
	}
}
