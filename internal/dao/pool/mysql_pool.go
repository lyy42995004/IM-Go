package pool

import (
	"fmt"

	"github.com/lyy42995004/IM-Go/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func init() {
	username := config.GetConfig().MySQL.User     // 账号
	password := config.GetConfig().MySQL.Password // 密码
	host := config.GetConfig().MySQL.Host         // 数据库地址
	port := config.GetConfig().MySQL.Port         // 数据库端口
	DBname := config.GetConfig().MySQL.Name       // 数据库名
	timeout := "10s"                              // 连接超时，10秒

	// 构建 DSN
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s",
		username, password, host, port, DBname, timeout)

	// 建立数据库连接
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}

	// 获取底层 SQL 连接实例，用于配置数据库连接池
	sqlDB, _ := db.DB()

	// 配置数据库连接池
	sqlDB.SetMaxOpenConns(100) // 设置数最大连接数
	sqlDB.SetMaxIdleConns(20)  // 最大允许的空闲连接数
}

// 返回数据库连接实例
func GetDB() *gorm.DB {
	return db
}
