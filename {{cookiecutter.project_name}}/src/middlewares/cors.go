package middlewares

import (
	"github.com/OVINC-CN/DevTemplateGo/src/configs"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CORS() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     configs.Config.AllowOrigins,
		AllowMethods:     configs.Config.AllowMethods,
		AllowHeaders:     configs.Config.AllowHeaders,
		ExposeHeaders:    configs.Config.ExposeHeaders,
		AllowCredentials: true,
	})
}
