package command

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"runtime"
	_ "the-video-project/backend/configs"
	"the-video-project/backend/internal/db"

	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type SeedCommand struct{}

func (cmd SeedCommand) Run() error {

	conn := db.Conn()

	info, err := ioutil.ReadDir(seedersFolder())
	if err != nil {
		return err
	}

	for _, v := range info {
		filepath := filepath.Join(seedersFolder(), v.Name())
		binary, err := ioutil.ReadFile(filepath)
		if err != nil {
			return err
		}
		tx := conn.Exec(string(binary))
		if tx.Error != nil {
			return tx.Error
		}
		fmt.Printf("%s (%d lines) were run\n", v.Name(), tx.RowsAffected)
	}
	return nil
}

func seedersFolder() string {
	_, b, _, _ := runtime.Caller(0)
	path := filepath.Join(filepath.Dir(b), "seeders/")
	return path
}
