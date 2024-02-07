package configs

import (
	"github.com/sirupsen/logrus"
	"time"
)

type configModel struct {
	AppCode   string
	AppSecret string
	Debug     bool
	LogLevel  logrus.Level
	serverConfigModel
	traceConfigModel
	corsConfigModel
	dbConfigModel
	sessionConfigModel
	redisConfigModel
}

type serverConfigModel struct {
	Addr           string
	RequestTimeout time.Duration
	TLSCert        string
	TLSKey         string
}

type corsConfigModel struct {
	AllowOrigins  []string
	AllowMethods  []string
	AllowHeaders  []string
	ExposeHeaders []string
}

type dbConfigModel struct {
	DBHost              string
	DBPort              string
	DBUser              string
	DBPassword          string
	DBName              string
	DBMaxConnections    int
	DBConnectionTimeOut time.Duration
	DBSlowThreshold     time.Duration
}

type sessionConfigModel struct {
	SessionCookieName     string
	SessionCookieAge      int
	SessionCookiePath     string
	SessionCookieDomain   string
	SessionCookieSecure   bool
	SessionCookieHttpOnly bool
}

type redisConfigModel struct {
	RedisHost           string
	RedisPort           string
	RedisPassword       string
	RedisDB             int
	RedisPrefix         string
	RedisMaxConnections int
}

type traceConfigModel struct {
	RUMID   string
	RUMHost string
}
