package main

import (
	"airun/brain-tracker/server/src/router"
)

func main() {
	router := router.GetRouter()
	router.Run(":8080")
}
