package configs

import (
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
)

func init() {
	// Get root path
	_, b, _, _ := runtime.Caller(0)
	envFilepath := filepath.Join(filepath.Dir(b), "..", ".env")

	if err := godotenv.Load(envFilepath); err != nil {
		log.Fatalln("Error loading .env file.")
	}
	if os.Getenv("APP_KEY") == "" {
		log.Fatalln("APP_KEY should be set.")
	}
}
