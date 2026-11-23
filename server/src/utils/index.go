package utils

import (
	"log"
	"os"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

func ResolveSuccessResponse(response string) gin.H {
	return gin.H{
		"message": response,
	}
}

func ResolveErrorResponse(error error) gin.H {
	return gin.H{
		"message": error.Error(),
	}
}

func GetEnv(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func LogError(err error) {
	log.Printf("Recieved Error %v , %s", err, debug.Stack())
}
