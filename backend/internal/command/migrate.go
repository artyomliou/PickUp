package command

import (
	"database/sql"
	"fmt"
	"path/filepath"
	"runtime"
	_ "the-video-project/backend/configs"
	"the-video-project/backend/internal/db"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type MigrateCommand struct {
	Subcommand string
}

func (cmd MigrateCommand) Run() error {

	conn, err := sql.Open("mysql", db.MysqlDsn())
	if err != nil {
		return err
	}

	driver, err := mysql.WithInstance(conn, &mysql.Config{})
	if err != nil {
		return err
	}
	m, err := migrate.NewWithDatabaseInstance(migrationsFolder(), "mysql", driver)
	if err != nil {
		return err
	}

	switch cmd.Subcommand {
	case "down":
		err := m.Down()
		if err != nil && err.Error() != "no change" {
			return err
		}
	default:
		err := m.Up()
		if err != nil && err.Error() != "no change" {
			return err
		}
	}
	return nil
}

func migrationsFolder() string {
	_, b, _, _ := runtime.Caller(0)
	path := filepath.Join(filepath.Dir(b), "migrations")
	return fmt.Sprintf("file://%s", path)
}
