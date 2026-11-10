package workers

import (
	"airun/brain-tracker/server/src/types"
	"context"
	"encoding/json"
	"log"
	"runtime/debug"
	"time"

	"github.com/gorilla/websocket"
)

var StreamSignalChannel = make(chan types.OutgoingClientSignalDTO[any], 256)
var ClientStreamSocket *websocket.Conn
var ClientStreamSocketActive bool

func InitiateClientStreamSocket(conn *websocket.Conn) {
	ClientStreamSocket = conn
}

func CloseClientStreamSocket() {
	err := ClientStreamSocket.Close()
	if err != nil {
		log.Printf("Error closing client stream socket: %v\n%s", err, debug.Stack())
	}
}

func InitiateStreamWorker(mainContext context.Context) {
	go func() {
		defer func() {
			err := recover()
			if err != nil {
				log.Printf("ERROR in stream worker %v\n%s", err, debug.Stack())
			}
		}()

		sendSignalToClient := func(signal *types.OutgoingClientSignalDTO[any]) {
			if !ClientStreamSocketActive {
				return
			}
			jsonSignal, err := json.Marshal(*signal)
			if err != nil {
				log.Printf("Error marshalling signal: %v\n%s", err, debug.Stack())
				return
			}
			ClientStreamSocket.SetWriteDeadline(time.Now().Add(10 * time.Second))
			err = ClientStreamSocket.WriteMessage(websocket.TextMessage, jsonSignal)
			if err != nil {
				log.Printf("Error sending signal to client: %v\n%s", err, debug.Stack())
				return
			}
		}

	mainLoop:
		for {
			select {
			case signal := <-StreamSignalChannel:
				sendSignalToClient(&signal)
			case <-mainContext.Done():
				break mainLoop
			}
		}
	}()
}
