package server

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/radovskyb/watcher"
)

func StartWatcher(opt *Options) error {
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
					fmt.Println("Modified file:", event.Name())
				case watcher.EventFileAdded:
					fmt.Println("Added file:", event.Name())
				case watcher.EventFileDeleted:
					fmt.Println("Deleted file:", event.Name())
				}
			case err := <-w.Error:
				log.Fatalln(err)
			}
		}
	}()

	if err := w.Add(opt.Only); err != nil {
		return err
	}

	if err := w.Start(time.Millisecond * 100); err != nil {
		return err
	}

	wg.Wait()

	return nil
}
