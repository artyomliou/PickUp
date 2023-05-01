package db

import (
	"path/filepath"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SqliteConn() (*gorm.DB, error) {
	return gorm.Open(sqlite.Open(SqliteDsn()), &gorm.Config{})
}

func SqliteDsn() string {
	path, _ := filepath.Abs("gorm.db")
	return path
}
