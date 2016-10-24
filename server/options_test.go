package server

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCfgFileName(t *testing.T) {
	assert.Equal(t, ".golivrc", cfgFileName, "should have the right name for the file")
}

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
	assert.Equal(t, "", o.HTTPURL, "default HTTPURL value")
	assert.Equal(t, "", o.WSURL, "default WSURL value")
}

func TestOptionsMount(t *testing.T) {
	o := NewOptions()

	o.Mount()

	assert.Equal(t, "http://127.0.0.1:1308", o.HTTPURL, "http - default mounted value")
	assert.Equal(t, "ws://127.0.0.1:1308/ws", o.WSURL, "ws - default mounted value")

	o.Host = "abc"
	o.Port = ":9876"

	o.Mount()

	assert.Equal(t, "http://abc:9876", o.HTTPURL, "http - custom mounted value - not secure")
	assert.Equal(t, "ws://abc:9876/ws", o.WSURL, "ws - custom mounted value - not secure")

	o.Host = "def"
	o.Port = ":1234"
	o.Secure = true

	o.Mount()

	assert.Equal(t, "https://def:1234", o.HTTPURL, "http - custom mounted value - secure")
	assert.Equal(t, "wss://def:1234/ws", o.WSURL, "ws - custom mounted value - secure")
}

func TestOptionsAssignBeingOverriddenByCli(t *testing.T) {
	opt1 := NewOptions()
	cli1 := NewOptions()
	file1 := NewOptions()

	cli1.Port = "abc"
	file1.Port = "123"

	opt1.Assign(*file1, *cli1)

	assert.Equal(t, "abc", cli1.Port, "should override the port from the file")

	opt2 := NewOptions()
	cli2 := NewOptions()
	file2 := NewOptions()

	cli2.Host = "https://abc.com"
	file2.Host = "yoyo://abc.??"

	opt2.Assign(*file2, *cli2)

	assert.Equal(t, "https://abc.com", cli2.Host, "should override the Host")
}

func TestOptionsAssignBeingOverriddenByFile(t *testing.T) {
	opt1 := NewOptions()
	cli1 := NewOptions()
	file1 := NewOptions()

	file1.Port = "123"

	opt1.Assign(*file1, *cli1)

	assert.Equal(t, "123", cli1.Port, "should override the port from the default values")

	opt2 := NewOptions()
	cli2 := NewOptions()
	file2 := NewOptions()

	file2.Host = "yoyo://abc.??"

	opt2.Assign(*file2, *cli2)

	assert.Equal(t, "yoyo://abc.??", cli2.Host, "should override the Host from the default values")
}

func TestOptionsAssignBeingAdded(t *testing.T) {
	opt1 := NewOptions()
	cli1 := NewOptions()
	file1 := NewOptions()

	cli1.Port = "abc"
	file1.Secure = true

	opt1.Assign(*file1, *cli1)

	assert.Equal(t, "abc", opt1.Port, "should keep the port as it was")
	assert.Equal(t, true, opt1.Secure, "should add the secure the the options")

	opt2 := NewOptions()
	cli2 := NewOptions()
	file2 := NewOptions()

	cli2.Host = "https://abc.com"
	file2.Only = "a,b,c"

	opt2.Assign(*file2, *cli2)

	assert.Equal(t, "https://abc.com", opt2.Host, "should keep the Host")
	assert.Equal(t, "a,b,c", opt2.Only, "should add only to the option")
}
