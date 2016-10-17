package server

func Start() error {
	if err := StartWatcher(); err != nil {
		return err
	}

	return nil
}
