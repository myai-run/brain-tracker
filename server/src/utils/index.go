package utils

import "github.com/gin-gonic/gin"

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
