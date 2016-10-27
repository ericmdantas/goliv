package server

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCfgFileName(t *testing.T) {
	assert.Equal(t, ".golivrc", cfgFileName, "should have the right name for the file")
}

func TestDefaultPort(t *testing.T) {
	assert.Equal(t, ":1308", defaultPort, "should have the right port value")
}

func TestDefaultHost(t *testing.T) {
	assert.Equal(t, "127.0.0.1", defaultHost, "should have the right host value")
}

func TestInlinePathSeparator(t *testing.T) {
	assert.Equal(t, ",", inlinePathSeparator, "should have the right path separator value")
}

func TestNewOptions(t *testing.T) {
	o := NewOptions()

	assert.Equal(t, ":1308", o.Port, "default proxy value")
	assert.Equal(t, "127.0.0.1", o.Host, "default host value")
	assert.Equal(t, false, o.Secure, "default secure value")
	assert.Equal(t, false, o.Quiet, "default quiet value")
	assert.Equal(t, false, o.NoBrowser, "default noBrowser value")
	assert.Equal(t, "", o.OnlyCLI, "default OnlyCLI value")
	assert.Equal(t, []string{}, o.Only, "default OnlyCLI value")
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

func TestOptionsParseURL(t *testing.T) {
	for _, v := range tableTestParseURL {
		o := NewOptions()

		o.Host = v.inHost
		o.Port = v.inPort
		o.Secure = v.inSecure

		o.Parse()

		assert.Equal(t, v.outHTTPURL, o.HTTPURL, v.descriptionHTTP)
		assert.Equal(t, v.outWSURL, o.WSURL, v.descriptionWS)
	}
}

func TestOptionsParseOnlyPaths(t *testing.T) {
	for _, v := range tableTestParseOnlyPaths {
		o := NewOptions()

		o.Only = v.inOnly
		o.OnlyCLI = v.inOnlyCLI

		o.Parse()

		assert.Equal(t, v.outOnly, o.Only, v.description)
	}
}

func TestOptionsAssignBeingTheDefaultValues(t *testing.T) {
	opt1 := NewOptions()
	default1 := *NewOptions()
	file1 := Options{}
	cli1 := Options{}

	opt1.Assign(default1, file1, cli1)

	assert.Equal(t, "127.0.0.1", default1.Host, "should have the default Host")
	assert.Equal(t, ":1308", default1.Port, "should have the default Port")
	assert.Equal(t, false, default1.Quiet, "should have the default Quiet")
	assert.Equal(t, false, default1.Secure, "should have the default Secure")
}

func TestOptionsAssignBeingOverriddenByCli(t *testing.T) {
	opt1 := NewOptions()
	default1 := *NewOptions()
	file1 := Options{}
	cli1 := Options{}

	cli1.Port = "abc"
	file1.Port = "123"

	opt1.Assign(default1, file1, cli1)

	assert.Equal(t, "abc", cli1.Port, "should override the port from the file")

	opt2 := NewOptions()
	default2 := *NewOptions()
	file2 := Options{}
	cli2 := Options{}

	cli2.Host = "https://abc.com"
	file2.Host = "yoyo://abc.??"

	opt2.Assign(default2, file2, cli2)

	assert.Equal(t, "https://abc.com", cli2.Host, "should override the Host")
}

func TestOptionsAssignBeingOverriddenByFile(t *testing.T) {
	opt1 := NewOptions()
	default1 := *NewOptions()
	file1 := Options{}
	cli1 := Options{}

	file1.Port = "123"

	opt1.Assign(default1, file1, cli1)

	assert.Equal(t, "123", opt1.Port, "should override the port from the default values")

	opt2 := NewOptions()
	default2 := *NewOptions()
	file2 := Options{}
	cli2 := Options{}

	file2.Host = "yoyo://abc.??"

	opt2.Assign(default2, file2, cli2)

	assert.Equal(t, "yoyo://abc.??", opt2.Host, "should override the Host from the default values")
}

func TestOptionsAssignBeingAdded(t *testing.T) {
	opt1 := NewOptions()
	default1 := *NewOptions()
	file1 := Options{}
	cli1 := Options{}

	cli1.Port = "abc"
	file1.Secure = true

	opt1.Assign(default1, file1, cli1)

	assert.Equal(t, "abc", opt1.Port, "should keep the port as it was")
	assert.Equal(t, true, opt1.Secure, "should add the secure the the options")

	opt2 := NewOptions()
	default2 := *NewOptions()
	file2 := Options{}
	cli2 := Options{}

	cli2.Host = "https://abc.com"
	file2.Only = []string{"a", "b", "c"}

	opt2.Assign(default2, file2, cli2)

	assert.Equal(t, "https://abc.com", opt2.Host, "should keep the Host")
	assert.Equal(t, []string{"a", "b", "c"}, opt2.Only, "should add Only to the option")

	opt3 := NewOptions()
	default3 := *NewOptions()
	file3 := Options{}
	cli3 := Options{}

	cli3.Host = "https://abc123.com"
	cli3.OnlyCLI = "x,y,z"

	opt3.Assign(default3, file3, cli3)

	assert.Equal(t, "https://abc123.com", opt3.Host, "should keep the Host")
	assert.Equal(t, "x,y,z", opt3.OnlyCLI, "should add OnlyCLI to the option")
	assert.Equal(t, []string{"x", "y", "z"}, opt3.Only, "should add Only to the option - already parsed")
}

var tableTestParseURL = []struct {
	inHost   string
	inPort   string
	inSecure bool

	outHTTPURL      string
	descriptionHTTP string

	outWSURL      string
	descriptionWS string
}{
	{
		inHost:          "",
		inPort:          "",
		inSecure:        false,
		outHTTPURL:      "http://127.0.0.1:1308",
		descriptionHTTP: "http - default parsed value",
		outWSURL:        "ws://127.0.0.1:1308/ws",
		descriptionWS:   "ws - default value",
	},
	{
		inHost:          "abc",
		inPort:          ":9876",
		inSecure:        false,
		outHTTPURL:      "http://abc:9876",
		descriptionHTTP: "http - custom parsed value - not secure",
		outWSURL:        "ws://abc:9876/ws",
		descriptionWS:   "ws - custom parsed value - not secure",
	},
	{
		inHost:          "def",
		inPort:          ":1234",
		inSecure:        true,
		outHTTPURL:      "https://def:1234",
		descriptionHTTP: "http - custom parsed value - secure",
		outWSURL:        "wss://def:1234/ws",
		descriptionWS:   "ws - default parsed value - secure",
	},
}

var tableTestParseOnlyPaths = []struct {
	inOnly    []string
	inOnlyCLI string

	outOnly []string

	description string
}{
	{
		inOnly:    []string{},
		inOnlyCLI: "",

		outOnly:     []string{},
		description: "only default value",
	},
	{
		inOnly:    []string{},
		inOnlyCLI: "a",

		outOnly:     []string{"a"},
		description: "single onlyCLI value being parsed into Only",
	},
	{
		inOnly:    []string{},
		inOnlyCLI: "a,b,c",

		outOnly:     []string{"a", "b", "c"},
		description: "multiple onlyCLI value being parsed into Only",
	},
	{
		inOnly:    []string{"x", "y", "z"},
		inOnlyCLI: "a,b,c",

		outOnly:     []string{"x", "y", "z"},
		description: "Only value being left as it is",
	},
}
