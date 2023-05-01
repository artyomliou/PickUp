package db

import (
	"gorm.io/gorm"
)

var conn *gorm.DB

// TODO return error
func Conn() *gorm.DB {
	if conn != nil {
		return conn
	}
	// driver := os.Getenv("DB_DRIVER")
	// if driver == "sqlite" {
	// 	c, _ := SqliteConn()
	// 	return c
	// }
	c, _ := MysqlConn()
	return c
}
