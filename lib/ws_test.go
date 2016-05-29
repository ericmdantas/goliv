package goliv

import "testing"

type connMock struct{}

func (cm connMock) Close() error { return nil }

var tableTestAddClients = []struct {
	in  Client
	out []Client
}{
	{
		in:  Client{Id: "a1", Conn: connMock{}},
		out: []Client{Client{Id: "a1", Conn: connMock{}}},
	},
	{
		in:  Client{Id: "a2", Conn: connMock{}},
		out: []Client{Client{Id: "a2", Conn: connMock{}}},
	},
}

func TestAddClients(t *testing.T) {
	cw0 := ClientWrapper{}

	cw0.Add(tableTestAddClients[0].in)

	if cw0.Clients[0] != tableTestAddClients[0].out[0] {
		t.Errorf("Expected %v to equal %v", cw0.Clients[0], tableTestAddClients[0].out)
	}

	cw1 := ClientWrapper{}

	cw1.Add(tableTestAddClients[1].in)

	if cw1.Clients[0] != tableTestAddClients[1].out[0] {
		t.Errorf("Expected %v to equal %v", cw1.Clients[0], tableTestAddClients[1].out)
	}
}

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
