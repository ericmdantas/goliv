package goliv

import (
	"testing"
)

func TestServerInstance(t *testing.T) {
	s := NewServer(&Options{
		Port: "123",
		Host: "abc.com",
	})

	if s.opts.Port != "123" {
		t.Errorf("Expected %s to equal %s", s.opts.Port, "123")
	}

	if s.opts.Host != "http://abc.com:123" {
		t.Errorf("Expected %s to equal %s", s.opts.Host, "http://abc.com:123")
	}
}
