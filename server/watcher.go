package server

import (
	"log"
	"sync"
	"time"

	"github.com/radovskyb/watcher"
)

func watchContent(cfg *Config, notifyChange func()) error {
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
					if !cfg.Quiet {
						log.Println("Created file:", event.Name())
					}

					notifyChange()
				case watcher.Write:
					if !cfg.Quiet {
						log.Println("Changed file:", event.Name())
					}

					notifyChange()
				case watcher.Remove:
					if !cfg.Quiet {
						log.Println("Removed file:", event.Name())
					}

					notifyChange()
				case watcher.Rename:
					if !cfg.Quiet {
						log.Println("Renamed file:", event.Name())
					}

					notifyChange()
				}
			case err := <-w.Error:
				log.Fatalln(err)
			}
		}
	}()

	for _, path := range cfg.Only {
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
