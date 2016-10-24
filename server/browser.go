package server

import (
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/skratchdot/open-golang/open"
)

var golivScript = `
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

func injectScript(o Options) (string, error) {
	file, err := ioutil.ReadFile(filepath.Join(o.PathIndex, "index.html"))

	if err != nil {
		return "", err
	}

	fileWithScript := strings.Replace(string(file), "</body>", golivScript+"</body>", -1)
	fileWithScript = strings.Replace(fileWithScript, "_WS_", o.WSURL, -1)

	return fileWithScript, nil
}

func openBrowser(opt Options) error {
	if opt.NoBrowser {
		return nil
	}

	return open.Start(opt.HTTPURL)
}
