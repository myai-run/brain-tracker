package router

import (
	"airun/brain-tracker/server/src/controllers"

	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	routerEngine := gin.Default()
	routerEngine.GET("/handle-signals", controllers.HandleIncomingSignals)
	return routerEngine
}
