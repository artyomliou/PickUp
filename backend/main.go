package main

import (
	"log"
	"os"
	"the-video-project/backend/http"
	"the-video-project/backend/models"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Panic(err)
	}

	if os.Getenv("APP_DEBUG") == "1" {
		models.DB().AutoMigrate(&models.User{})
		(&models.Seeder{}).Users()
	}

	router := http.SetupRouter()
	router.Run(":5000")
}
