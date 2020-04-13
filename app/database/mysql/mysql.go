package mysql

import (
	"fmt"
	"gin-app-start/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var MysqlClient *gorm.DB

func Init() *gorm.DB {
	var err error
	MysqlClient, err = gorm.Open(config.DRIVER, config.MysqlUrl)
	MysqlClient.LogMode(true)
	if err != nil {
		fmt.Println("mysql connection error: ", err)
		panic(err)
	}
	fmt.Println("mysql connection open to: ", config.MysqlUrl)
	return MysqlClient
}
