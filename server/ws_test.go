package server

import (
	"testing"

	"github.com/gorilla/websocket"

	"github.com/stretchr/testify/assert"
)

func TestNewWS(t *testing.T) {
	ws := NewWS()

	assert.Equal(t, []*websocket.Conn{}, ws.clients, "should start with an empty client slice")
}

func TestWSScript(t *testing.T) {
	script := `
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

	assert.Contains(t, script, WSScript, "should have the right info for the script")
}

func TestOnConnect(t *testing.T) {
	ws := NewWS()

	c := websocket.Conn{}

	ws.OnConnect(&c, func() {

	})
}

func TestBroadcastReload(t *testing.T) {
	ws := NewWS()
	conn1 := websocket.Conn{}
	conn2 := websocket.Conn{}

	ws.clients = append(ws.clients, &conn1)
	ws.clients = append(ws.clients, &conn2)

	ws.BroadcastReload()
}
