package model

import (
	"goblog/pkg/logger"

	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	// GORM 的 MSYQL 数据库驱动导入
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DB gorm.DB 对象
var DB *gorm.DB

// ConnectDB 初始化模型
func ConnectDB() *gorm.DB {

	var err error

	// 设置数据库连接信息
	config := mysql.Config{
		User:                 "root",
		Passwd:               "secret",
		Addr:                 "127.0.0.1:3306",
		Net:                  "tcp",
		DBName:               "goblog",
		AllowNativePasswords: true,
	}

	// 准备数据库连接池
	DB, err = gorm.Open("mysql", config.FormatDSN())
	logger.LogError(err)

	return DB
}
