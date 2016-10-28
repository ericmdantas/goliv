package server

import (
	"strings"
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

func TestInjectScript(t *testing.T) {
	o := NewOptions()

	o.Parse()
	o.indexHTMLContent = []byte(`
	!!!<body></body>???
`)

	indexWithScript := `
	!!!<body><div id="goliv-container" style="display: none">
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
		</div></body>???
`

	indexWithScriptInjected, err := injectScript(o)

	if err != nil {
		assert.Fail(t, "error injecting the script")
	}

	r := strings.NewReplacer("\n", "", "\t", "")

	indexWithScript = r.Replace(indexWithScript)
	indexWithScriptInjected = r.Replace(indexWithScriptInjected)

	assert.Equal(t, indexWithScript, indexWithScriptInjected, "should have the right html injected")
}
