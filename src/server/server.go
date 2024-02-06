package main

import (
	"github.com/OrenZhang/GoDevTest/src/configs"
	"github.com/OrenZhang/GoDevTest/src/db"
	"github.com/OrenZhang/GoDevTest/src/services/account"
	"github.com/OrenZhang/GoDevTest/src/utils"
)

func startServer() {
	// init log
	utils.InitLogger(configs.Config.Debug, configs.Config.LogLevel)
	// init db
	db.InitDBConnection(configs.DBConfig.Host, configs.DBConfig.Port, configs.DBConfig.User, configs.DBConfig.Password, configs.DBConfig.Name)
	migrate()
	// init redis
	db.InitRedisConnection(configs.RedisConfig.Host, configs.RedisConfig.Port, configs.RedisConfig.Password, configs.RedisConfig.DB)
	// init gin
	engine := setupRouter()
	if err := engine.Run(configs.Config.Port); err != nil {
		utils.Logger.Infof("[ServerStartFailed] %s", err)
		panic(err)
	}
}

func migrate() {
	err := db.DB.AutoMigrate(account.User{}, account.UserSession{})
	if err != nil {
		utils.Logger.Errorf("[MigrateDBFailed] %s", err)
	} else {
		utils.Logger.Infof("[MigrateDBSuccess] %T", err)
	}
}
