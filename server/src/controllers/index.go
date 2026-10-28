package controllers

import (
	"airun/brain-tracker/server/src/configs"
	"airun/brain-tracker/server/src/models"
	"airun/brain-tracker/server/src/repository"
	"airun/brain-tracker/server/src/types"
	"airun/brain-tracker/server/src/utils"
	"encoding/json"
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

		parsedSignal := map[string]interface{}{}
		err = json.Unmarshal(message, &parsedSignal)
		if err != nil {
			log.Printf("Unmarshal Error: %v", err)
			continue
		}
		signalType := utils.GetSignalDataType(parsedSignal)
		if signalType == types.UNKNOWN {
			continue
		}
		repoValue := models.SignalData{
			Values: parsedSignal[string(signalType)].([]interface{}),
			Sid:    parsedSignal["sid"].(string),
			Time:   parsedSignal["time"].(float64),
			Type:   signalType,
		}
		err = repository.CreateSignalData(&repoValue)
		if err != nil {
			log.Printf("Signal Creation Failed: %v", err)
		}

		log.Printf("Inserted: %s", message)

		if err := wsConnection.WriteMessage(messageType, message); err != nil {
			log.Printf("Write error: %v", err)
			break
		}
	}
}
