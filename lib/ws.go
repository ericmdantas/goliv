package goliv

const (
	WsContainerId    = "#goliv-container"
	WsInjectedScript = `
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
)
