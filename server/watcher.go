package server

import (
	"log"
	"strings"
	"sync"
	"time"

	"github.com/radovskyb/watcher"
)

type ContentWatcher struct {
	options           *Options
	watchablePathsRaw string
	WatchablePaths    []string
}

func (cw *ContentWatcher) Watch(notifyChange func()) error {
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
				case watcher.EventFileModified:
					if !cw.options.Quiet {
						log.Println("Modified file:", event.Name())
					}

					notifyChange()
				case watcher.EventFileAdded:
					if !cw.options.Quiet {
						log.Println("Added file:", event.Name())
					}

					notifyChange()
				case watcher.EventFileDeleted:
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

func NewContentWatcher(opt *Options) *ContentWatcher {
	rawPath := opt.Only
	splitPaths := strings.Split(rawPath, ",")

	return &ContentWatcher{
		options:           opt,
		watchablePathsRaw: opt.Only,
		WatchablePaths:    splitPaths,
	}
}
