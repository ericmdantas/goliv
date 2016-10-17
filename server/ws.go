package server

import (
	"github.com/gorilla/websocket"
)

type WS struct {
	conn *websocket.Conn
}
