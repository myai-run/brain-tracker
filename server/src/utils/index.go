package utils

import (
	"os"

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
