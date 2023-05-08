package db

import (
	"fmt"
	"os"
	_ "pick-up/backend/configs"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func MysqlConn() (*gorm.DB, error) {
	return gorm.Open(mysql.Open(MysqlDsn()), &gorm.Config{})
}

func MysqlDsn() string {
	user := os.Getenv("MYSQL_USER")
	pass := os.Getenv("MYSQL_PASSWORD")
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	dbname := os.Getenv("MYSQL_DATABASE")

	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port, dbname)
}
