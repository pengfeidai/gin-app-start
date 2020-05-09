package mysql

import (
	"gin-app-start/app/common"
	"gin-app-start/app/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func Init() {
	var err error
	logger := common.Logger
	config := config.Conf.Mysql
	url := config.User + ":" + config.Password + "@(" + config.Path + ")/" + config.Database + "?" + config.Config
	db, err := gorm.Open(config.Driver, url)
	if err != nil {
		logger.Error("mysql connection error: ", err)
		panic(err)
	}
	// db.DB().SetMaxIdleConns(config.MaxIdleConns)
	// db.DB().SetMaxOpenConns(config.MaxOpenConns)
	// 全局禁用表名复数
	db.SingularTable(true)
	db.LogMode(config.Log)

	// 注册表
	db.AutoMigrate()

	logger.Info("mysql connection open to: ", url)
	DB = db
}
