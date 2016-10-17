package server

func Start(opt *Options) error {
	opt.Mount()

	OpenBrowser(opt)

	if err := StartWatcher(opt); err != nil {
		return err
	}

	return nil
}
