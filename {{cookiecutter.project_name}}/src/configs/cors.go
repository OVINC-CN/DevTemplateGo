package configs

import (
	"github.com/OVINC-CN/DevTemplateGo/src/utils"
	"strings"
)

type CORSConfigModel struct {
	AllowOrigins  []string
	AllowMethods  []string
	AllowHeaders  []string
	ExposeHeaders []string
}

var CORSConfig = CORSConfigModel{
	AllowOrigins:  strings.Split(utils.GetEnv("CORS_ALLOW_ORIGINS", ""), ";"),
	AllowMethods:  strings.Split(utils.GetEnv("CORS_ALLOW_METHODS", "*"), ";"),
	AllowHeaders:  strings.Split(utils.GetEnv("CORS_ALLOW_HEADERS", "Content-Type"), ";"),
	ExposeHeaders: strings.Split(utils.GetEnv("CORS_EXPOSE_HEADERS", "Content-Length"), ";"),
}
