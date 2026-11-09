package main

import (
	"airun/brain-tracker/server/src/configs"
	"airun/brain-tracker/server/src/router"
	"airun/brain-tracker/server/src/workers"
	"context"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	mainContext, cancelMainContext := context.WithCancel(context.Background())
	defer cancelMainContext()

	godotenv.Load(".env")
	err := configs.InitiateDatabaseConnection()
	if err != nil {
		log.Panic("Database Connection Failed!")
		return
	}

	workers.InitiateDBWorker(mainContext)
	workers.InitiateStreamWorker(mainContext)

	router := router.GetRouter()
	router.Run(":8080")
}
