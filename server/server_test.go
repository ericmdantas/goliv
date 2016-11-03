package server

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServerConstant(t *testing.T) {
	t.Run("reloadEvent", func(t *testing.T) {
		assert.Equal(t, "reload", reloadEvent, "should have the right value for the event")
	})

	t.Run("intervalFileCheck", func(t *testing.T) {
		assert.Equal(t, 1000, intervalFileCheck, "should have the right value for the interval")
	})
}
