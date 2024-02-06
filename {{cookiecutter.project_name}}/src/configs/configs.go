package configs

import (
	"fmt"
	"github.com/OVINC-CN/DevTemplateGo/src/utils"
	"github.com/sirupsen/logrus"
	"os"
	"strconv"
	"strings"
	"time"
)

var Config configModel

func InitConfig() {
	loadEnv()
	Config = configModel{
		AppCode:   os.Getenv("APP_CODE"),
		AppSecret: os.Getenv("APP_SECRET"),
		Debug:     utils.StrToBool(utils.GetEnv("DEBUG", "false")),
		LogLevel:  logrus.InfoLevel,
		serverConfigModel: serverConfigModel{
			Addr:           utils.GetEnv("SERVER_ADDR", ":8000"),
			RequestTimeout: time.Duration(utils.StrToInt(utils.GetEnv("REQUEST_TIMEOUT", "10"))) * time.Second,
			TLSCert:        utils.GetEnv("TLS_CERT", ""),
			TLSKey:         utils.GetEnv("TLS_KEY", ""),
		},
		ConfigModel: ConfigModel{
			SessionCookieName:     buildSessionCookieName(),
			SessionCookieAge:      utils.StrToInt(utils.GetEnv("SESSION_COOKIE_AGE", strconv.Itoa(60*60*24*7))),
			SessionCookiePath:     utils.GetEnv("SESSION_COOKIE_PATH", "/"),
			SessionCookieDomain:   os.Getenv("SESSION_COOKIE_DOMAIN"),
			SessionCookieSecure:   utils.StrToBool(utils.GetEnv("SESSION_COOKIE_SECURE", "false")),
			SessionCookieHttpOnly: utils.StrToBool(utils.GetEnv("SESSION_COOKIE_HTTP_ONLY", "false")),
		},
		corsConfigModel: corsConfigModel{
			AllowOrigins:  strings.Split(utils.GetEnv("CORS_ALLOW_ORIGINS", ""), ";"),
			AllowMethods:  strings.Split(utils.GetEnv("CORS_ALLOW_METHODS", "*"), ";"),
			AllowHeaders:  strings.Split(utils.GetEnv("CORS_ALLOW_HEADERS", "Content-Type"), ";"),
			ExposeHeaders: strings.Split(utils.GetEnv("CORS_EXPOSE_HEADERS", "Content-Length"), ";"),
		},
		dbConfigModel: dbConfigModel{
			DBHost:              utils.GetEnv("DB_HOST", "127.0.0.1"),
			DBPort:              utils.GetEnv("DB_PORT", "3306"),
			DBUser:              utils.GetEnv("DB_USER", ""),
			DBPassword:          utils.GetEnv("DB_PASSWORD", ""),
			DBName:              utils.GetEnv("DB_NAME", ""),
			DBMaxConnections:    utils.StrToInt(utils.GetEnv("DB_MAX_CONNECTIONS", "10")),
			DBConnectionTimeOut: time.Duration(utils.StrToInt(utils.GetEnv("DB_CONNECTION_TIMEOUT", strconv.Itoa(60*60)))) * time.Second,
			DBSlowThreshold:     time.Duration(utils.StrToInt(utils.GetEnv("DB_SLOW_THRESHOLD", strconv.Itoa(100)))) * time.Millisecond,
		},
		redisConfigModel: redisConfigModel{
			RedisHost:           utils.GetEnv("REDIS_HOST", "127.0.0.1"),
			RedisPort:           utils.GetEnv("REDIS_PORT", "6379"),
			RedisPassword:       utils.GetEnv("REDIS_PASSWORD", ""),
			RedisDB:             utils.StrToInt(utils.GetEnv("REDIS_DB", "0")),
			RedisPrefix:         utils.GetEnv("REDIS_PREFIX", ""),
			RedisMaxConnections: utils.StrToInt(utils.GetEnv("REDIS_MAX_CONNECTIONS", "10")),
		},
		traceConfigModel: traceConfigModel{
			RUMID:   utils.GetEnv("RUM_ID", ""),
			RUMHost: utils.GetEnv("RUM_HOST", "https://rumt-zh.com"),
		},
	}
}

func buildSessionCookieName() (cookieName string) {
	var devFlag string
	if Config.Debug {
		devFlag = "-dev"
	} else {
		devFlag = ""
	}
	cookieName = fmt.Sprintf("%s-session-id%s", Config.AppCode, devFlag)
	return
}
