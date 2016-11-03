package server

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBrowserConstant(t *testing.T) {
	t.Run("golivscript", func(t *testing.T) {
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
	})
}

func TestInjectScript(t *testing.T) {
	t.Run("inject", func(t *testing.T) {
		for _, v := range tableTestInjectScript {
			cfg := NewConfig()

			cfg.parse()
			cfg.indexHTMLContent = v.inIndex
			indexWithScriptInjected, err := injectScript(cfg)

			if err != nil {
				assert.Fail(t, "error injecting the script")
			}

			r := strings.NewReplacer("\n", "", "\t", "")

			indexWithScript := r.Replace(v.outIndex)
			indexWithScriptInjected = r.Replace(indexWithScriptInjected)

			assert.Equal(t, indexWithScript, indexWithScriptInjected, v.description)
		}
	})
}

var tableTestInjectScript = []struct {
	inIndex []byte

	outIndex string

	description string
}{
	{
		inIndex: []byte("!!!<body></body>???"),
		outIndex: `
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
		`,
		description: "should have the right html injected - symbols",
	},
	{
		inIndex: []byte("123<body></body>456"),
		outIndex: `
			123<body><div id="goliv-container" style="display: none">
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
				</div></body>456
		`,
		description: "should have the right html injected - numbers",
	},
	{
		inIndex: []byte("<body></body>"),
		outIndex: `
			<body><div id="goliv-container" style="display: none">
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
				</div></body>
		`,
		description: "should have the right html injected - only body",
	},
	{
		inIndex:     []byte(""),
		outIndex:    "",
		description: "should not append anything, html is empty",
	},
	{
		inIndex:     []byte("<html></html>"),
		outIndex:    "<html></html>",
		description: "should not append anything, html has no body",
	},
	{
		inIndex:     []byte("<body>"),
		outIndex:    "<body>",
		description: "should not append anything, html doesn't have a closing body",
	},
	{
		inIndex: []byte("</body>"),
		outIndex: `<div id="goliv-container" style="display: none">
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
				</div></body>`,
		description: "should append the script, html has a closing body",
	},
}
