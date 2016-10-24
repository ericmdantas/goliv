package server

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCfgFileName(t *testing.T) {
	assert.Equal(t, ".golivrc", cfgFileName, "should have the right name for the file")
}
