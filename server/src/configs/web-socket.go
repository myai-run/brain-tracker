package configs

import "github.com/gorilla/websocket"

func GetWebSocketInitiator() *websocket.Upgrader {
	var upgrader = websocket.Upgrader{}
	return &upgrader
}
