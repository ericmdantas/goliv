package server

import (
	"github.com/skratchdot/open-golang/open"
)

func OpenBrowser(opt *Options) error {
	return open.Start(opt.URL)
}
