package server

import (
	"io/ioutil"
	"strings"
)

var golivScript = `
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

func InjectScript(o *Options) (string, error) {
	file, err := ioutil.ReadFile(o.PathIndex + "index.html")

	if err != nil {
		return "", err
	}

	fileWithScript := strings.Replace(string(file), "</body>", golivScript+"</body>", -1)

	return fileWithScript, nil
}
