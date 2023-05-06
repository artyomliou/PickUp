package db

import (
	"log"
	"os"

	"gorm.io/gorm"
)

var conn *gorm.DB

func Conn() *gorm.DB {
	if conn != nil {
		return conn
	}

	var err error
	var c *gorm.DB

	driver := os.Getenv("DB_DRIVER")
	switch driver {
	case "sqlite":
		c, err = SqliteConn()

	default:
		c, err = MysqlConn()
	}
	if err != nil {
		log.Panic(err)
	}
	conn = c
	return conn
}
