package server

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReloadEventConstant(t *testing.T) {
	assert.Equal(t, "reload", reloadEvent, "should have the right value for the event")
}

func TestIntervalFileCheckConstant(t *testing.T) {
	assert.Equal(t, 1000, intervalFileCheck, "should have the right value for the interval")
}
