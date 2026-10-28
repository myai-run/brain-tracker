package main

import (
	"airun/brain-tracker/server/src/configs"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
		return
	}

	err = configs.InitiateDatabaseConnection()
	if err != nil {
		log.Fatal(err)
		return
	}

	err = configs.Migrate()
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Print("Migration Done Successfully!")
}
