package db

import (
	"fmt"
	"github.com/OVINC-CN/DevTemplateGo/src/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDBConnection(host, port, user, password, name string) {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4", user, password, host, port, name)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		utils.Logger.Errorf("[InitDBConnectionFailed] %s", err)
	} else {
		utils.Logger.Infof("[InitDBConnectionSuccess] %T", DB)
	}
}
