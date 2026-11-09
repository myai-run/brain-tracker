package workers

import (
	"airun/brain-tracker/server/src/types"
	"context"
	"log"

	"github.com/gorilla/websocket"
)

var StreamSignalChannel = make(chan types.OutgoingClientSignalDTO[any])
var ClientStreamSocket *websocket.Conn

func InitiateClientStreamSignal(conn *websocket.Conn) {
	ClientStreamSocket = conn
}

func InitiateStreamWorker(context context.Context) {
	go func() {
		err := recover()
		if err != nil {
			log.Printf("ERROR in stream worker %v", err)
		}

		sendSignalToClient := func(signal *types.OutgoingClientSignalDTO[any]) types.OutgoingClientSignalDTO[any] {
			// WORKING ON HERE
			err := ClientStreamSocket.WriteMessage(0, *signal)
			if err != nil {
				log.Printf("Error sending signal to client: %v", err)
			}
		}

		for {
			select {
			case signal := <-StreamSignalChannel:
				sendSignalToClient(&signal)
			case <-context.Done():
				break
			}
		}
	}()
}
