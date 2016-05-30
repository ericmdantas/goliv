package goliv

import "golang.org/x/net/websocket"

const (
	WsContainerId    = "#goliv-container"
	WsInjectedScript = `
    <div id="goliv-container" style="display:none">
      <script>
        ;(function() {
          var _protocol = /^https/.test(location.protocol) ? 'wss' : 'ws';
          var _port = 1234;

          var _ws = new WebSocket(_protocol + "//" + location.host + "/ws")

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

type Closer interface {
	Close() error
}

type Client struct {
	Id   string
	Conn Closer
}

type ClientWrapper struct {
	Clients []Client
}

func (cw *ClientWrapper) Connected(conn *websocket.Conn) {
	cl := Client{
		Id:   "1",
		Conn: conn,
	}

	cw.Add(cl)
}

func (cw *ClientWrapper) Add(c Client) {
	cw.Clients = append(cw.Clients, c)
}

func (cw *ClientWrapper) Remove(id string) {
	for i, c := range cw.Clients {
		if c.Id == id {
			cw.Clients = append(cw.Clients[:i], cw.Clients[i+1:]...)
		}
	}
}
