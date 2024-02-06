package configs

import "github.com/OVINC-CN/DevTemplateGo/src/utils"

type redisConfigModel struct {
	Host           string
	Port           string
	Password       string
	DB             int
	Prefix         string
	MaxConnections int
}

var RedisConfig = redisConfigModel{
	Host:           utils.GetEnv("REDIS_HOST", "127.0.0.1"),
	Port:           utils.GetEnv("REDIS_PORT", "6379"),
	Password:       utils.GetEnv("REDIS_PASSWORD", ""),
	DB:             utils.StrToInt(utils.GetEnv("REDIS_DB", "0")),
	Prefix:         utils.GetEnv("REDIS_PREFIX", ""),
	MaxConnections: utils.StrToInt(utils.GetEnv("REDIS_MAX_CONNECTIONS", "10")),
}
