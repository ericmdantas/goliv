package server

import (
	"github.com/gorilla/websocket"
)

var WSScript = `
	<div id="goliv-container" style="display: none">
		;(function() {
			"use strict";
			
			var ws = new WebSocket("ws://127.0.0.1:986/ws");
			
			ws.onmessage = function(ev) {
				if (ev.data === "reload") {
					window.location.reload();
				}
			};
		}());
	</div>
`

type WSSender interface {
	WriteMessage(messageType int, payload []byte) error
}

type WS struct {
	clients []*WSSender
}

func (ws *WS) OnConnect(conn WSSender, cb func()) {
	ws.clients = append(ws.clients, conn)

	cb()
}

func (ws *WS) BroadcastReload() {
	for _, client := range ws.clients {
		client.WriteMessage(websocket.TextMessage, []byte("reload"))
	}
}

func NewWS() *WS {
	return &WS{
		clients: []*websocket.Conn{},
	}
}
