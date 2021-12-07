package db

import (
	settings "github.com/ewenliu/gin-example/7.Gorm/env"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var _db *gorm.DB

func SetupDB() {
	var err error
	dbConnMap := settings.Env["DB"].(map[string]string)
	// 这里Open返回的是线程池，并非单独的mysql连接线程
	_db, err = gorm.Open("mysql", dbConnMap["default"])
	if err != nil {
		panic("连接数据库失败, error: " + err.Error())
	}
	sqlDB := _db.DB()
	sqlDB.SetMaxOpenConns(200)
	sqlDB.SetMaxIdleConns(20)
}

func GetDB() *gorm.DB {
	return _db
}
