package db

import (
	"log"

	"gorm.io/gorm"
)

var conn *gorm.DB

func Conn() *gorm.DB {
	if conn != nil {
		return conn
	}
	// driver := os.Getenv("DB_DRIVER")
	// if driver == "sqlite" {
	// 	c, _ := SqliteConn()
	// 	return c
	// }
	c, err := MysqlConn()
	if err != nil {
		log.Panic(err)
	}
	return c
}
