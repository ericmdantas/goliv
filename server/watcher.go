package server

import (
	"log"
	"sync"
	"time"

	"github.com/radovskyb/watcher"
)

func watchContent(opt *Options, notifyChange func()) error {
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
					if !opt.Quiet {
						log.Println("Created file:", event.Name())
					}

					notifyChange()
				case watcher.Write:
					if !opt.Quiet {
						log.Println("Changed file:", event.Name())
					}

					notifyChange()
				case watcher.Remove:
					if !opt.Quiet {
						log.Println("Removed file:", event.Name())
					}

					notifyChange()
				case watcher.Rename:
					if !opt.Quiet {
						log.Println("Renamed file:", event.Name())
					}

					notifyChange()
				}
			case err := <-w.Error:
				log.Fatalln(err)
			}
		}
	}()

	for _, path := range opt.Only {
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
