package server

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewContentWatcherSimplePath(t *testing.T) {
	opts := &Options{
		Only: "abc",
	}

	cw := NewContentWatcher(opts)

	assert.Equal(t, opts, cw.options, "should have the same options")
	assert.Equal(t, "abc", cw.watchablePathsRaw, "should have the same path raw")
	assert.Equal(t, []string{"abc"}, cw.WatchablePaths, "should have the same path parsed")
}

func TestNewContentWatcherComplexPaths(t *testing.T) {
	opts := &Options{
		Only: "a,b,c,d,e,f",
	}

	cw := NewContentWatcher(opts)

	assert.Equal(t, opts, cw.options, "should have the same options")
	assert.Equal(t, "a,b,c,d,e,f", cw.watchablePathsRaw, "should have the same path raw")
	assert.Equal(t, []string{"a", "b", "c", "d", "e", "f"}, cw.WatchablePaths, "should have the same path parsed")
}
