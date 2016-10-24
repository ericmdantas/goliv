package server

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGolivScript(t *testing.T) {
	script := `
	<div id="goliv-container" style="display: none">
		<script>	
			;(function() {
				"use strict";
				
				var ws = new WebSocket("_WS_");
				
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
