package middlewares

import (
	"github.com/OrenZhang/GoDevTest/src/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func InitLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 初始化请求ID
		requestID := utils.GenerateUniqID()
		// 初始化Logger
		logEntry := utils.Logger.WithFields(logrus.Fields{"request_id": requestID})
		c.Set("logEntry", logEntry)
		c.Next()
	}
}