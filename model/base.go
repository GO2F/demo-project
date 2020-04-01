package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // 基于mysql引擎
	"rank/config"
)

// DB 数据库实例
var DB *gorm.DB = nil

// InitDb 初始化数据库实例
func InitDb() (err error) {
	// parseTime=true 不加这个无法解析时间字段, 会导致数据绑定失败, 无数据返回
	db, err := gorm.Open("mysql", config.Db.User+":"+config.Db.Password+"@("+config.Db.Host+":"+config.Db.Port+")/"+config.Db.Database+"?charset=utf8mb4&loc=Local&parseTime=true")
	if err != nil {
		fmt.Println("数据库启动异常%S", err)
		return nil
	}
	DB = db
	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)
	return nil
}
