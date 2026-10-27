package controllers

import (
	"airun/brain-tracker/server/src/configs"
	"log"

	"github.com/gin-gonic/gin"
)

func HandleIncomingSignals(requestContext *gin.Context) {
	wsConnection, err := configs.GetWebSocketInitiator().Upgrade(
		requestContext.Writer,
		requestContext.Request,
		nil,
	)
	if err != nil {
		log.Printf("WebSocket upgrade error: %v", err)
		return
	}
	defer wsConnection.Close()

	for {
		messageType, message, err := wsConnection.ReadMessage()
		if err != nil {
			log.Printf("Read error: %v", err)
			break
		}
		log.Printf("Received: %s", message)

		if err := wsConnection.WriteMessage(messageType, message); err != nil {
			log.Printf("Write error: %v", err)
			break
		}
	}
}
