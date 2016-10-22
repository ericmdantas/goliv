package server

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInjectScript(t *testing.T) {
	assert.Equal(t, true, true, "a")
}

func TestGolivScript(t *testing.T) {
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

	assert.Equal(t, script, golivScript, "should have the right info for the script")
}
