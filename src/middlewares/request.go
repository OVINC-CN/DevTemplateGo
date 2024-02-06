package middlewares

import (
	"github.com/OVINC-CN/DevTemplateGo/src/services/account"
	"github.com/OVINC-CN/DevTemplateGo/src/utils"
	"github.com/gin-gonic/gin"
	"time"
)

func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 初始化请求时间
		t := time.Now()
		// 执行
		c.Next()
		// 记录请求耗时
		duration := time.Since(t).Milliseconds()
		// 记录用户
		username := account.GetContextUser(c).Username
		if username == "" {
			username = "-"
		}
		// 记录请求日志
		utils.ContextInfof(c, "[RequestLog] %s %s %s %d %d", username, c.Request.Method, c.Request.URL, duration, c.Writer.Status())
	}
}
