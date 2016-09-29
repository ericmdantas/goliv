package goliv

import (
	"time"

	"github.com/radovskyb/watcher"
)

type SourceWatcher struct {
	w *watcher.Watcher
}

func NewWatcher() *SourceWatcher {
	return &SourceWatcher{
		w: watcher.New(),
	}
}

func (s *SourceWatcher) Add(info string) error {
	return s.w.Add(info)
}

func (s *SourceWatcher) Remove(info string) error {
	return s.w.Remove(info)
}

func (s *SourceWatcher) Start(t time.Duration) error {
	return s.w.Start(t)
}
