package workers

import (
	"airun/brain-tracker/server/src/types"
	"airun/brain-tracker/server/src/utils"
	"context"
	"encoding/json"
	"errors"

	"github.com/gorilla/websocket"
)

var StreamChannel = make(chan types.OutgoingClientSignalDTO[any])
var StreamChannelActive bool = false

type StreamWorkerOptions struct {
	MainCtx      context.Context
	WsConnection *websocket.Conn
}

func InitiateStreamWorker(options StreamWorkerOptions) {
	go func() {
		defer func() {
			err := recover()
			if err != nil {
				utils.LogError(err.(error))
			}
		}()

		for {
			_, message, err := options.WsConnection.ReadMessage()
			unMarshaledSignal := types.IncomingClientSignalDTO[any]{}
			if err != nil {
				utils.LogError(
					errors.New("Error While Reading Socket"),
				)
				continue
			}
			err = json.Unmarshal(message, &unMarshaledSignal)
			if err != nil {
				utils.LogError(
					errors.New("Validating Incoming Signal Failed"),
				)
				continue
			}
			if unMarshaledSignal.Message == types.START_STREAMING {
				StreamChannelActive = true
				break
			} else if unMarshaledSignal.Message == types.END_STREAMING {
				StreamChannelActive = false
				break
			}
		}

		sendToClient := func(outgoingSignal types.OutgoingClientSignalDTO[any]) {
			if !StreamChannelActive {
				utils.LogError(
					errors.New("Stream Channel Activation Required Before Sending Any Message"),
				)
				return
			}

			stringifiedMessage, err := json.Marshal(outgoingSignal)
			if err != nil {
				utils.LogError(
					errors.New("Outgoing Signal Marshaling Failed"),
				)
				return
			}
			err = options.WsConnection.WriteMessage(websocket.TextMessage, stringifiedMessage)
			if err != nil {
				utils.LogError(
					errors.New("Outgoing Signal Write Failed"),
				)
				return
			}
		}

		select {
		case signal := <-StreamChannel:
			sendToClient(signal)
		case <-options.MainCtx.Done():
			options.WsConnection.Close()
		}
	}()
}
