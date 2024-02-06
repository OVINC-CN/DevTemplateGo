package configs

import (
	"github.com/OrenZhang/GoDevTest/src/utils"
	"github.com/sirupsen/logrus"
	"os"
)

type configModel struct {
	Debug     bool
	LogLevel  logrus.Level
	Port      string
	AppCode   string
	AppSecret string
}

var Config = configModel{
	Debug:     utils.StrToBool(utils.GetEnv("DEBUG", "false")),
	LogLevel:  logrus.InfoLevel,
	Port:      utils.GetEnv("PORT", ":8000"),
	AppCode:   os.Getenv("APP_CODE"),
	AppSecret: os.Getenv("APP_SECRET"),
}
