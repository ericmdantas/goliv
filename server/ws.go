package server

import (
	_ "github.com/gorilla/websocket"
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
