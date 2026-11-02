package router

import (
	"airun/brain-tracker/server/src/controllers"

	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	routerEngine := gin.Default()
	signalRoutes := routerEngine.Group("signal")
	signalRoutes.GET("handle-resource", controllers.HandleResourceSignals)
	signalRoutes.GET("handle-client", controllers.HandleClientSignals)
	return routerEngine
}
