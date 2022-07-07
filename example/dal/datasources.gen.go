package dal

import (
	"fmt"
	

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db1 = ConnectDB("root:passwd@(127.0.0.1:3306)/ai_iotman?charset=utf8mb4&parseTime=True").Debug()
	
)

func ConnectDB(dsn string) (db *gorm.DB) {
	conn, err := gorm.Open(mysql.Open(dsn))

	if err != nil {
		panic(fmt.Errorf("connect db fail: %w", err))
	}

	return conn
}

func CreateDbConn(modelName string) *gorm.DB {
	return db1
}
