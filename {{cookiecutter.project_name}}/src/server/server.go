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
	var err error
	// init config
	configs.InitConfig()
	// init log
	utils.Logger = utils.InitLogger(
		configs.Config.Debug,
		configs.Config.LogLevel,
	)
	utils.DbLogger = utils.InitDBLogger(
		configs.Config.Debug,
		configs.Config.LogLevel,
		configs.Config.DBSlowThreshold,
	)
	// init db
	db.InitDBConnection(
		configs.Config.DBHost,
		configs.Config.DBPort,
		configs.Config.DBUser,
		configs.Config.DBPassword,
		configs.Config.DBName,
		configs.Config.DBMaxConnections,
		configs.Config.DBConnectionTimeOut,
		&gorm.Config{
			Logger: utils.DbLogger,
		},
	)
	migrate()
	// init redis
	db.InitRedisConnection(
		configs.Config.RedisHost,
		configs.Config.RedisPort,
		configs.Config.RedisPassword,
		configs.Config.RedisDB,
		configs.Config.RedisMaxConnections,
	)
	// init cpu
	threads := runtime.NumCPU()
	runtime.GOMAXPROCS(threads)
	utils.Logger.Infof("[InitCPUSuccess] Runs on %d CPUs", threads)
	// init gin
	engine := setupRouter()
	if configs.Config.TLSCert != "" {
		err = engine.RunTLS(configs.Config.Addr, configs.Config.TLSCert, configs.Config.TLSKey)
	} else {
		err = engine.Run(configs.Config.Addr)
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
