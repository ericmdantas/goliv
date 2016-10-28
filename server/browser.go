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

func injectScript(o *Options) (string, error) {
	fileWithScript := strings.Replace(string(o.indexHTML), "</body>", golivScript+"</body>", -1)
	fileWithScript = strings.Replace(fileWithScript, "_WS_", o.WSURL, -1)

	return fileWithScript, nil
}

func openBrowser(opt *Options) error {
	if opt.NoBrowser {
		return nil
	}

	return open.Start(opt.HTTPURL)
}
