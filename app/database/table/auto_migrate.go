package table

import (
	"gin-app-start/app/common"
	"gin-app-start/app/database/mysql"
	"gin-app-start/app/model"
)

var logger = common.Logger

func Init() {
	db := mysql.DB
	db.AutoMigrate(&model.User{})
	logger.Info("AutoMigrate tables success.")
}
