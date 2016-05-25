package goliv

import "testing"

func TestConstantWsContainerId(t *testing.T) {
	id := "#goliv-container"

	if WsContainerId != id {
		t.Errorf("Expected %s to equal %s", WsContainerId, id)
	}
}

func TestConstantWsInjectedHtml(t *testing.T) {
	r := `
    <div id="goliv-container" style="display:none">
      <script>
        ;(function() {
          var _protocol = /^https/.test(location.protocol) ? 'wss' : 'ws';
          var _port = 1234;

          var _ws = new WebSocket(_protocol + "//" + location.host)

          _ws.onmessage = function(ev) {
          if (ev.data === "reload") {
            location.reload();
          }
        }
        }());
      </script>
    </div>
  `

	if WsInjectedScript != r {
		t.Errorf("Expected %s to equal %s", WsInjectedScript, r)
	}
}
