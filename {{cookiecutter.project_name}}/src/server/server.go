package main

import (
	"github.com/OVINC-CN/DevTemplateGo/src/configs"
	"github.com/OVINC-CN/DevTemplateGo/src/db"
	"github.com/OVINC-CN/DevTemplateGo/src/services/account"
	"github.com/OVINC-CN/DevTemplateGo/src/utils"
	"gorm.io/gorm"
	"runtime"
)

func startServer() {
	// init log
	utils.Logger = utils.InitLogger(
		configs.Config.Debug,
		configs.Config.LogLevel,
	)
	utils.DbLogger = utils.InitDBLogger(
		configs.Config.Debug,
		configs.Config.LogLevel,
		configs.DBConfig.SlowThreshold,
	)
	// init db
	db.InitDBConnection(
		configs.DBConfig.Host,
		configs.DBConfig.Port,
		configs.DBConfig.User,
		configs.DBConfig.Password,
		configs.DBConfig.Name,
		configs.DBConfig.MaxConnections,
		configs.DBConfig.ConnectionTimeOut,
		&gorm.Config{
			Logger: utils.DbLogger,
		},
	)
	migrate()
	// init redis
	db.InitRedisConnection(
		configs.RedisConfig.Host,
		configs.RedisConfig.Port,
		configs.RedisConfig.Password,
		configs.RedisConfig.DB,
		configs.RedisConfig.MaxConnections,
	)
	// init cpu
	threads := runtime.NumCPU()
	runtime.GOMAXPROCS(threads)
	utils.Logger.Infof("[InitCPUSuccess] Runs on %d CPUs", threads)
	// init gin
	engine := setupRouter()
	var err error
	if configs.Config.TLSCert != "" {
		err = engine.RunTLS(configs.Config.Port, configs.Config.TLSCert, configs.Config.TLSKey)
	} else {
		err = engine.Run(configs.Config.Port)
	}
	if err != nil {
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
