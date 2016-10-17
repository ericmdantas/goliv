package server

import (
	"github.com/skratchdot/open-golang/open"
)

func OpenBrowser() error {
	return open.Start("https://google.com")
}
