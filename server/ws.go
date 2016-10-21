package server

import (
	"github.com/gorilla/websocket"
)

const (
	evReload = "reload"
)

var WSScript = `
	<div id="goliv-container" style="display: none">
		<script>	
			;(function() {
				"use strict";
				
				var ws = new WebSocket("ws://127.0.0.1:1308/ws");
				
				ws.onmessage = function(ev) {
					if (ev.data === "reload") {
						window.location.reload();
					}
				};
			}());
		</script>
	</div>
`

type WS struct {
	clients []*websocket.Conn
}

func (ws *WS) OnConnect(conn *websocket.Conn, cb func()) {
	ws.clients = append(ws.clients, conn)

	cb()
}

func (ws *WS) BroadcastReload() {
	for _, client := range ws.clients {
		client.WriteMessage(websocket.TextMessage, []byte(evReload))
	}
}

func NewWS() *WS {
	return &WS{
		clients: []*websocket.Conn{},
	}
}
