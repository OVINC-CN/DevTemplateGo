package db

import (
	"fmt"
	"github.com/OVINC-CN/DevTemplateGo/src/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var DB *gorm.DB

func InitDBConnection(host, port, user, password, name string, maxConnections int, connectionTimeout time.Duration, config *gorm.Config) {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4", user, password, host, port, name)
	DB, err = gorm.Open(mysql.Open(dsn), config)
	sqlDB, err := DB.DB()
	sqlDB.SetMaxOpenConns(maxConnections)
	sqlDB.SetConnMaxLifetime(connectionTimeout)
	if err != nil {
		utils.Logger.Errorf("[InitDBConnectionFailed] %s", err)
	} else {
		utils.Logger.Infof("[InitDBConnectionSuccess] %T", DB)
	}
}
