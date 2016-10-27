package server

import (
	"log"
	"sync"
	"time"

	"github.com/radovskyb/watcher"
)

type contentWatcher struct {
	options        *Options
	WatchablePaths []string
}

func (cw contentWatcher) Watch(notifyChange func()) error {
	var wg sync.WaitGroup

	w := watcher.New()
	w.SetMaxEvents(1)

	wg.Add(1)

	go func() {
		defer wg.Done()

		for {
			select {
			case event := <-w.Event:
				switch event.Op {
				case watcher.Create:
					if !cw.options.Quiet {
						log.Println("Created file:", event.Name())
					}

					notifyChange()
				case watcher.Write:
					if !cw.options.Quiet {
						log.Println("Changed file:", event.Name())
					}

					notifyChange()
				case watcher.Remove:
					if !cw.options.Quiet {
						log.Println("Removed file:", event.Name())
					}

					notifyChange()

				case watcher.Rename:
					if !cw.options.Quiet {
						log.Println("Renamed file:", event.Name())
					}

					notifyChange()

				}
			case err := <-w.Error:
				log.Fatalln(err)
			}
		}
	}()

	for _, path := range cw.WatchablePaths {
		if err := w.Add(path); err != nil {
			return err
		}
	}

	if err := w.Start(time.Millisecond * 100); err != nil {
		return err
	}

	wg.Wait()

	return nil
}

func newContentWatcher(opt *Options) *contentWatcher {
	return &contentWatcher{
		options:        opt,
		WatchablePaths: opt.Only,
	}
}
