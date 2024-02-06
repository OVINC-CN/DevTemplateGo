package configs

import "github.com/OVINC-CN/DevTemplateGo/src/utils"

type dbConfigModel struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

var DBConfig = dbConfigModel{
	Host:     utils.GetEnv("DB_HOST", "127.0.0.1"),
	Port:     utils.GetEnv("DB_PORT", "3306"),
	User:     utils.GetEnv("DB_USER", ""),
	Password: utils.GetEnv("DB_PASSWORD", ""),
	Name:     utils.GetEnv("DB_NAME", ""),
}
