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

func HandleClientSignals(requestContext *gin.Context) {
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

		parsedSignal := types.ClientSignalDataDTO[any]{}
		err = json.Unmarshal(message, &parsedSignal)
		if err != nil {
			log.Printf("Unmarshal Error: %v", err)
			continue
		}
		if !utils.IsValidClientSignal(&parsedSignal) {
			continue
		}

		// stream start
		if parsedSignal.Type == types.START_STREAMING {

		}

		// manage streaming
		if parsedSignal.Type == types.STREAMING {
			// get data from repository and send to client
			err := wsConnection.WriteMessage(messageType, message)
			if err != nil {
				continue
			}
			var signals = []models.SignalData{}
			var params = types.GetSingalDataParamsDTO{
				// Offset: ,
			}
			signals, err := repository.GetSignalData()
		}

		// manage stream end
		if parsedSignal.Type == types.END_STREAMING {
			log.Print("End Stream command recieved. Closing now ...")
			err := wsConnection.Close()
			if err != nil {
				log.Print("Stream Close Failed.")
				log.Print(err)
			} else {
				log.Print("Stream Closed")
			}
			break
		}
	}
}

func HandleSignals(requestContext *gin.Context) {
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
