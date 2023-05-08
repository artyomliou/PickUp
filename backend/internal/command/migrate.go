package command

import (
	_ "pick-up/backend/configs"
	"pick-up/backend/internal/db"
	"pick-up/backend/internal/models"

	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type MigrateCommand struct {
	Subcommand string
}

func (cmd MigrateCommand) Run() error {

	migrator := db.Conn().Migrator()

	list := []interface{}{
		models.User{},
		models.Store{},
		models.Product{},
		models.Category{},
		models.SelectQuestion{},
		models.SelectOption{},
		models.CustomQuestion{},
		models.Cart{},
		models.CartItem{},
		models.Order{},
	}

	switch cmd.Subcommand {
	case "down":
		for i := range list {
			m := list[len(list)-1-i]
			if migrator.HasTable(m) {
				if err := migrator.DropTable(m); err != nil {
					return err
				}
			}
		}
	default:
		for _, m := range list {
			if !migrator.HasTable(m) {
				if err := migrator.AutoMigrate(m); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

// func migrationsFolder() string {
// 	_, b, _, _ := runtime.Caller(0)
// 	path := filepath.Join(filepath.Dir(b), "migrations")
// 	return fmt.Sprintf("file://%s", path)
// }
