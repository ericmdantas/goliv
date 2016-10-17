package server

func Start(opt *Options) error {
	if err := StartWatcher(); err != nil {
		return err
	}

	return nil
}
