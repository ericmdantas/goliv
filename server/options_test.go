package server

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewOptions(t *testing.T) {
	o := NewOptions()

	assert.Equal(t, ":1308", o.Port, "default proxy value")
	assert.Equal(t, "127.0.0.1", o.Host, "default host value")
	assert.Equal(t, false, o.Secure, "default secure value")
	assert.Equal(t, false, o.Quiet, "default quiet value")
	assert.Equal(t, false, o.NoBrowser, "default noBrowser value")
	assert.Equal(t, "", o.Only, "default only value")
	assert.Equal(t, "", o.Ignore, "default ignore value")
	assert.Equal(t, "", o.PathIndex, "default pathIndex value")
	assert.Equal(t, false, o.Proxy, "default proxy value")
	assert.Equal(t, "", o.ProxyTarget, "default proxyTarget value")
	assert.Equal(t, "", o.ProxyWhen, "default proxyWhen value")
	assert.Equal(t, "", o.Root, "default root value")
	assert.Equal(t, "", o.Static, "default static value")
	assert.Equal(t, "", o.URL, "default url value")
}

func TestOptionsMount(t *testing.T) {
	o := NewOptions()

	o.Mount()

	assert.Equal(t, "http://127.0.0.1:1308", o.URL, "default mounted value")

	o.Host = "abc"
	o.Port = ":9876"

	o.Mount()

	assert.Equal(t, "http://abc:9876", o.URL, "custom mounted value - not secure")

	o.Host = "def"
	o.Port = ":1234"
	o.Secure = true

	o.Mount()

	assert.Equal(t, "https://def:1234", o.URL, "custom mounted value - secure")
}
