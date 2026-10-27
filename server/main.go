package main

import (
	"airun/brain-tracker/server/src/router"
	"log"
)

func main() {
	err := router.GetRouter()
	if err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
