package config

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

const defaultDSN = "root:123456@tcp(127.0.0.1:3306)/zhenlizhiquan?charset=utf8mb4&parseTime=True&loc=Local"

func InitDB() {
	db, err := gorm.Open(mysql.Open(envOr("DB_DSN", defaultDSN)), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
	})
	if err != nil {
		log.Fatal("连接数据库失败:", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("获取数据库连接失败:", err)
	}
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(time.Hour)

	if err := sqlDB.Ping(); err != nil {
		log.Fatal("数据库 Ping 失败:", err)
	}

	DB = db
	log.Println("数据库连接成功")
}

func envOr(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
