package server

import "github.com/skratchdot/open-golang/open"

func OpenBrowser(opt *Options) error {
	if opt.NoBrowser {
		return nil
	}

	return open.Start(opt.URL)
}
