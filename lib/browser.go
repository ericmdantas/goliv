package goliv

type BrowserOpener interface {
	OpenBrowser() error
}

func OpenBrowser(bo BrowserOpener) error {
	return bo.OpenBrowser()
}
