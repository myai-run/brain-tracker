package main

import (
	"airun/brain-tracker/server/src/configs"
	"airun/brain-tracker/server/src/router"
	"airun/brain-tracker/server/src/workers"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	err := configs.InitiateDatabaseConnection()
	if err != nil {
		log.Panic("Database Connection Failed!")
		return
	}

	workers.InitiateDBWorker()

	router := router.GetRouter()
	router.Run(":8080")
}
