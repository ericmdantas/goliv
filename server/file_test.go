package server

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewIndexFile(t *testing.T) {
	t.Run("default", func(t *testing.T) {
		cfg := NewConfig()

		cfg.Root = "abc"
		cfg.HTTP2 = true
		cfg.Only = []string{"x", "y", "z"}

		f := newIndexFile(cfg)

		assert.Equal(t, cfg, f.cfg, "should have the same cfg in indexFile.cfg")
	})
}
