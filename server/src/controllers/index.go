package controllers

import (
	"airun/brain-tracker/server/src/configs"
	"airun/brain-tracker/server/src/models"
	"airun/brain-tracker/server/src/types"
	"airun/brain-tracker/server/src/utils"
	"airun/brain-tracker/server/src/workers"
	"encoding/json"
	"log"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

func HandleResourceSignals(requestContext *gin.Context) {
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
		streamValue := types.OutgoingClientSignalDTO[any]{
			Message: types.LIVE_STREAMING,
			Type:    signalType,
			Data:    parsedSignal[string(signalType)].([]interface{}),
		}

		// save data while streaming it
		workers.DbWriteChannel <- repoValue
		// live stream
		workers.StreamSignalChannel <- streamValue

		if err := wsConnection.WriteMessage(messageType, message); err != nil {
			log.Printf("Write error: %v", err)
			break
		}
	}
}

func HandleClientSignals(requestContext *gin.Context) {
	wsConnection, err := configs.GetWebSocketInitiator().Upgrade(
		requestContext.Writer,
		requestContext.Request,
		nil,
	)
	if err != nil {
		log.Printf("WebSocket upgrade error: %v\n%s", err, debug.Stack())
		return
	}
	defer wsConnection.Close()

	streamWorkerOptions := workers.StreamWorkerOptions{
		WsConnection: wsConnection,
		MainCtx:      requestContext.Request.Context(),
	}
	workers.InitiateStreamWorker(streamWorkerOptions)
}
