package configs

import (
	"github.com/OVINC-CN/DevTemplateGo/src/utils"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

type configModel struct {
	AppCode        string
	AppSecret      string
	Debug          bool
	LogLevel       logrus.Level
	Port           string
	RequestTimeout time.Duration
	TLSCert        string
	TLSKey         string
	RUMID          string
	RUMHost        string
}

var Config = configModel{
	AppCode:        os.Getenv("APP_CODE"),
	AppSecret:      os.Getenv("APP_SECRET"),
	Debug:          utils.StrToBool(utils.GetEnv("DEBUG", "false")),
	LogLevel:       logrus.InfoLevel,
	Port:           utils.GetEnv("PORT", ":8000"),
	RequestTimeout: time.Duration(utils.StrToInt(utils.GetEnv("REQUEST_TIMEOUT", "10"))) * time.Second,
	TLSCert:        utils.GetEnv("TLS_CERT", ""),
	TLSKey:         utils.GetEnv("TLS_KEY", ""),
	RUMID:          utils.GetEnv("RUM_ID", ""),
	RUMHost:        utils.GetEnv("RUM_HOST", "https://rumt-zh.com"),
}
