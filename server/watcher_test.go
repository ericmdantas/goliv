package server

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewContentWatcherSimplePath(t *testing.T) {
	opts := &Options{
		Only: []string{"abc"},
	}

	cw := newContentWatcher(opts)

	assert.Equal(t, opts, cw.options, "should have the same options")
	assert.Equal(t, []string{"abc"}, cw.WatchablePaths, "should have the same path parsed")
}

func TestNewContentWatcherComplexPaths(t *testing.T) {
	opts := &Options{
		Only: []string{"a", "b", "c", "d", "e", "f"},
	}

	cw := newContentWatcher(opts)

	assert.Equal(t, opts, cw.options, "should have the same options")
	assert.Equal(t, []string{"a", "b", "c", "d", "e", "f"}, cw.WatchablePaths, "should have the same path parsed")
}
