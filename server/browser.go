package server

import (
	"strings"

	"github.com/skratchdot/open-golang/open"
)

const golivScript = `
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

func injectScript(cfg *Config) (string, error) {
	fileWithScript := strings.Replace(string(cfg.indexHTMLContent), "</body>", golivScript+"</body>", -1)
	fileWithScript = strings.Replace(fileWithScript, "_WS_", cfg.WSURL, -1)

	return fileWithScript, nil
}

func openBrowser(cfg *Config) error {
	if cfg.NoBrowser {
		return nil
	}

	return open.Start(cfg.HTTPURL)
}
