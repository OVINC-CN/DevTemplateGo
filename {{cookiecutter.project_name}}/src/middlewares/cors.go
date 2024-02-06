package middlewares

import (
	"github.com/OVINC-CN/DevTemplateGo/src/configs"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CORS() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     configs.CORSConfig.AllowOrigins,
		AllowMethods:     configs.CORSConfig.AllowMethods,
		AllowHeaders:     configs.CORSConfig.AllowHeaders,
		ExposeHeaders:    configs.CORSConfig.ExposeHeaders,
		AllowCredentials: true,
	})
}
