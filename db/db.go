package db

import (
	"Chise1/imooc-product/ErrorInfo"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

var dbConn *gorm.DB

func InitDB(dbName string) {
	if dbConn == nil {
		//配置连接池
		db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
		if err != nil {

			panic("Failed to connect database.")
		}
		dbConn = db
		sqlDB, err := dbConn.DB()
		sqlDB.SetConnMaxLifetime(time.Hour)
		sqlDB.SetMaxIdleConns(5)
		sqlDB.SetMaxOpenConns(100)
	}
}

func GetDB() (*gorm.DB, error) {
	if dbConn != nil {
		return dbConn, nil
	}
	return dbConn, ErrorInfo.DbConnectionError
}
