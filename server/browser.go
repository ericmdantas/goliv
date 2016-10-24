package server

import "github.com/skratchdot/open-golang/open"

func openBrowser(opt Options) error {
	if opt.NoBrowser {
		return nil
	}

	return open.Start(opt.HTTPURL)
}
