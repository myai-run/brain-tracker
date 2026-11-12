package controllers

import (
	"airun/brain-tracker/server/src/configs"
	"airun/brain-tracker/server/src/models"
	"airun/brain-tracker/server/src/repository"
	"airun/brain-tracker/server/src/types"
	"airun/brain-tracker/server/src/utils"
	"airun/brain-tracker/server/src/workers"
	"encoding/json"
	"log"
	"runtime/debug"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
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

		workers.DbWriteChannel <- repoValue

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
	workers.InitiateClientStreamSocket(wsConnection)
	defer workers.CloseClientStreamSocket()

	limit := 100
	for {
		_, message, err := wsConnection.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				// handle unexpected close event
				workers.ClientStreamSocketActive = false
			}

			log.Printf("Read error: %v", err)
			break
		}

		parsedSignal := types.IncomingClientSignalDTO[any]{}
		err = json.Unmarshal(message, &parsedSignal)
		if err != nil {
			log.Printf("Unmarshal Error: %v", err)
			continue
		}

		if !utils.IsValidClientSignal(&parsedSignal) {
			continue
		}

		if parsedSignal.Type == types.END_STREAMING {
			// handle expected close signal
			workers.ClientStreamSocketActive = false
			log.Print("Recieved END_STREAM signal, closing")
			break
		}

		if parsedSignal.Type == types.START_STREAMING {
			workers.ClientStreamSocketActive = true
		}

		if workers.ClientStreamSocketActive {
			params := types.GetSingalDataParamsDTO{
				Limit:         &limit,
				RecievedAfter: parsedSignal.LastStreamRecieved,
			}
			signals, err := repository.GetSignalData(params)
			if err != nil {
				log.Printf("Error getting signal data: %v\n%s", err, debug.Stack())
				continue
			}
			for _, signal := range signals {
				workers.StreamSignalChannel <- utils.SignalDataToOutgoingClientSignalDTO(signal)
			}
		}
	}
}
