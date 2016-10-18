package server

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type conn struct {
}

func (c conn) WriteMessage(t int, pay []byte) error {
	return nil
}

func TestNewWS(t *testing.T) {
	ws := NewWS()

	assert.Equal(t, []WSSender, ws.clients, "should start with an empty client slice")
}

func TestWSScript(t *testing.T) {
	script := `
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

	assert.Contains(t, script, WSScript, "should have the right info for the script")
}

func TestOnConnect(t *testing.T) {
	ws := NewWS()

	c := conn{}

	ws.OnConnect(&c, func() {

	})
}

func TestBroadcastReload(t *testing.T) {
	ws := NewWS()
	conn1 := conn{}
	conn2 := conn{}

	ws.clients = append(ws.clients, &conn1)
	ws.clients = append(ws.clients, &conn2)

	ws.BroadcastReload()
}
