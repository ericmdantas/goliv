package server

import (
	"log"
	"strings"
	"sync"
	"time"

	"github.com/radovskyb/watcher"
)

type contentWatcher struct {
	options           Options
	watchablePathsRaw string
	WatchablePaths    []string
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
				switch event.EventType {
				case watcher.Modify:
					if !cw.options.Quiet {
						log.Println("Modified file:", event.Name())
					}

					notifyChange()
				case watcher.Add:
					if !cw.options.Quiet {
						log.Println("Added file:", event.Name())
					}

					notifyChange()
				case watcher.Remove:
					if !cw.options.Quiet {
						log.Println("Deleted file:", event.Name())
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

func newContentWatcher(opt Options) *contentWatcher {
	rawPath := opt.Only
	splitPaths := strings.Split(rawPath, ",")

	return &contentWatcher{
		options:           opt,
		watchablePathsRaw: opt.Only,
		WatchablePaths:    splitPaths,
	}
}
