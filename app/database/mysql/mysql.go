package mysql

import (
	"gin-app-start/config"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var Pool *gorm.DB

func Init() *gorm.DB {
	var err error
	db, err := gorm.Open(config.DRIVER, config.MysqlUrl)
	if err != nil {
		log.Println("mysql connection error: ", err)
		panic(err)
	}
	// db.DB().SetMaxIdleConns(config.MaxIdleConns)
	// db.DB().SetMaxOpenConns(config.MaxOpenConns)
	// 全局禁用表名复数
	db.SingularTable(true)
	db.LogMode(true)
	log.Println("mysql connection open to: ", config.MysqlUrl)
	Pool = db
	return Pool
}
