package configs

import (
	"fmt"
	"github.com/OrenZhang/GoDevTest/src/utils"
	"os"
	"strconv"
)

type sessionConfigModel struct {
	SessionCookieName     string
	SessionCookieAge      int
	SessionCookiePath     string
	SessionCookieDomain   string
	SessionCookieSecure   bool
	SessionCookieHttpOnly bool
}

var SessionConfig = sessionConfigModel{
	SessionCookieName:     buildSessionCookieName(),
	SessionCookieAge:      int(utils.StrToInt(utils.GetEnv("SESSION_COOKIE_AGE", strconv.Itoa(60*60*24*7)))),
	SessionCookiePath:     utils.GetEnv("SESSION_COOKIE_PATH", "/"),
	SessionCookieDomain:   os.Getenv("SESSION_COOKIE_DOMAIN"),
	SessionCookieSecure:   utils.StrToBool(utils.GetEnv("SESSION_COOKIE_SECURE", "false")),
	SessionCookieHttpOnly: utils.StrToBool(utils.GetEnv("SESSION_COOKIE_HTTP_ONLY", "false")),
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
