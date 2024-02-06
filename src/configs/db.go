package configs

import (
	"github.com/OVINC-CN/DevTemplateGo/src/utils"
	"strconv"
	"time"
)

type dbConfigModel struct {
	Host              string
	Port              string
	User              string
	Password          string
	Name              string
	MaxConnections    int
	ConnectionTimeOut time.Duration
	SlowThreshold     time.Duration
}

var DBConfig = dbConfigModel{
	Host:              utils.GetEnv("DB_HOST", "127.0.0.1"),
	Port:              utils.GetEnv("DB_PORT", "3306"),
	User:              utils.GetEnv("DB_USER", ""),
	Password:          utils.GetEnv("DB_PASSWORD", ""),
	Name:              utils.GetEnv("DB_NAME", ""),
	MaxConnections:    utils.StrToInt(utils.GetEnv("DB_MAX_CONNECTIONS", "10")),
	ConnectionTimeOut: time.Duration(utils.StrToInt(utils.GetEnv("DB_CONNECTION_TIMEOUT", strconv.Itoa(60*60)))) * time.Second,
	SlowThreshold:     time.Duration(utils.StrToInt(utils.GetEnv("DB_SLOW_THRESHOLD", strconv.Itoa(100)))) * time.Millisecond,
}
