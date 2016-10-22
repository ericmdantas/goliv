package server

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReloadEvent(t *testing.T) {
	assert.Equal(t, "reload", reloadEvent, "should have the right value for the event")
}
