package goliv

import (
	"testing"
)

type bo struct {
	called bool
}

func (b *bo) OpenBrowser() error {
	b.called = true

	return nil
}

func TestOpenBrowser(t *testing.T) {
	b := bo{false}

	OpenBrowser(&b)

	if b.called != true {
		t.Errorf("Expected %v to equal %v", b.called, true)
	}
}
