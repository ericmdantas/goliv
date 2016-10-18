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
		;(function() {
			"use strict";
			
			var ws = new WebSocket("ws://127.0.0.1:986/abc");
			
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
