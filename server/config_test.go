package server

import (
	"errors"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

type myIndexFileMock struct {
	info string
	err  error
}

func (m myIndexFileMock) readIndexHTML() ([]byte, error) {
	return []byte(m.info), m.err
}

func TestCfgFileNameConstant(t *testing.T) {
	assert.Equal(t, ".golivrc", cfgFileName, "should have the right name for the file")
}

func TestDefaultPortConstant(t *testing.T) {
	assert.Equal(t, ":1308", defaultPort, "should have the right port value")
}

func TestDefaultHostConstant(t *testing.T) {
	assert.Equal(t, "127.0.0.1", defaultHost, "should have the right host value")
}

func TestInlinePathSeparatorConstant(t *testing.T) {
	assert.Equal(t, ",", inlinePathSeparator, "should have the right path separator value")
}

func TestNameIndexHTMLFileConstant(t *testing.T) {
	assert.Equal(t, "index.html", nameIndexHTMLFile, "should have the right name for the index file")
}

func TestNewConfig(t *testing.T) {
	cfg := NewConfig()

	assert.Equal(t, ":1308", cfg.Port, "default proxy value")
	assert.Equal(t, "127.0.0.1", cfg.Host, "default host value")
	assert.Equal(t, false, cfg.HTTP2, "default http2 value")
	assert.Equal(t, false, cfg.Quiet, "default quiet value")
	assert.Equal(t, false, cfg.NoBrowser, "default noBrowser value")
	assert.Equal(t, "", cfg.OnlyCLI, "default OnlyCLI value")
	assert.Equal(t, []string{}, cfg.Only, "default OnlyCLI value")
	assert.Equal(t, []string{}, cfg.Ignore, "default ignore value")
	assert.Equal(t, "", cfg.PathIndex, "default pathIndex value")
	assert.Equal(t, false, cfg.Proxy, "default proxy value")
	assert.Equal(t, "", cfg.ProxyTarget, "default proxyTarget value")
	assert.Equal(t, "", cfg.ProxyWhen, "default proxyWhen value")
	assert.Equal(t, "", cfg.Root, "default root value")
	assert.Equal(t, "", cfg.Static, "default static value")
	assert.Equal(t, "", cfg.HTTPURL, "default HTTPURL value")
	assert.Equal(t, "", cfg.WSURL, "default WSURL value")
	assert.Equal(t, "", cfg.indexHTMLPath, "default index.html path")
	assert.Equal(t, []byte{}, cfg.indexHTMLContent, "default index.html content value")
	assert.Nil(t, cfg.indexHTMLFile)
}

func TestConfigParseURL(t *testing.T) {
	for _, v := range tableTestParseURL {
		cfg := NewConfig()

		cfg.Host = v.inHost
		cfg.Port = v.inPort
		cfg.HTTP2 = v.inSecure

		cfg.parse()

		assert.Equal(t, v.outHTTPURL, cfg.HTTPURL, v.descriptionHTTP)
		assert.Equal(t, v.outWSURL, cfg.WSURL, v.descriptionWS)
	}
}

func TestConfigParseOnlyPaths(t *testing.T) {
	for _, v := range tableTestParseOnlyPaths {
		cfg := NewConfig()

		cfg.Only = v.inOnly
		cfg.OnlyCLI = v.inOnlyCLI
		cfg.Root = v.inRoot

		cfg.parse()

		assert.Equal(t, v.outOnly, cfg.Only, v.description)
	}
}

func TestConfigParseIgnorePaths(t *testing.T) {
	for _, v := range tableTestParseIgnorePaths {
		cfg := NewConfig()

		cfg.Ignore = v.inIgnoreFile
		cfg.IgnoreCLI = v.inIgnoreCLI
		cfg.Root = v.inRoot

		cfg.parse()

		assert.Equal(t, v.outIgnore, cfg.Ignore, v.description)
	}
}

func TestConfigParseIndexHTMLPathInfo(t *testing.T) {
	for _, v := range tableTestParseIndexHTMLPathInfo {
		cfg := NewConfig()

		cfg.Root = v.inRoot
		cfg.PathIndex = v.inPathIndex

		cfg.parse()

		assert.Equal(t, v.outIndexHTMLPath, cfg.indexHTMLPath, v.description)
	}
}

func TestConfigAssignBeingTheDefaultValues(t *testing.T) {
	opt1 := NewConfig()
	default1 := *NewConfig()
	file1 := Config{}
	cli1 := Config{}

	opt1.assign(default1, file1, cli1)

	assert.Equal(t, "127.0.0.1", default1.Host, "should have the default Host")
	assert.Equal(t, ":1308", default1.Port, "should have the default Port")
	assert.Equal(t, false, default1.Quiet, "should have the default Quiet")
	assert.Equal(t, false, default1.HTTP2, "should have the default HTTP2")
}

func TestConfigAssignBeingOverriddenByCli(t *testing.T) {
	opt1 := NewConfig()
	default1 := *NewConfig()
	file1 := Config{}
	cli1 := Config{}

	cli1.Port = "abc"
	file1.Port = "123"

	opt1.assign(default1, file1, cli1)

	assert.Equal(t, "abc", cli1.Port, "should override the port from the file")

	opt2 := NewConfig()
	default2 := *NewConfig()
	file2 := Config{}
	cli2 := Config{}

	cli2.Host = "https://abc.com"
	file2.Host = "yoyo://abc.??"

	opt2.assign(default2, file2, cli2)

	assert.Equal(t, "https://abc.com", cli2.Host, "should override the Host")
}

func TestConfigAssignBeingOverriddenByFile(t *testing.T) {
	opt1 := NewConfig()
	default1 := *NewConfig()
	file1 := Config{}
	cli1 := Config{}

	file1.Port = "123"

	opt1.assign(default1, file1, cli1)

	assert.Equal(t, "123", opt1.Port, "should override the port from the default values")

	opt2 := NewConfig()
	default2 := *NewConfig()
	file2 := Config{}
	cli2 := Config{}

	file2.Host = "yoyo://abc.??"

	opt2.assign(default2, file2, cli2)

	assert.Equal(t, "yoyo://abc.??", opt2.Host, "should override the Host from the default values")
}

func TestConfigAssignBeingAdded(t *testing.T) {
	opt1 := NewConfig()
	default1 := *NewConfig()
	file1 := Config{}
	cli1 := Config{}

	cli1.Port = "abc"
	file1.HTTP2 = true

	opt1.assign(default1, file1, cli1)

	assert.Equal(t, "abc", opt1.Port, "should keep the port as it was")
	assert.Equal(t, true, opt1.HTTP2, "should add the http2 the the config")

	opt2 := NewConfig()
	default2 := *NewConfig()
	file2 := Config{}
	cli2 := Config{}

	cli2.Host = "https://abc.com"
	file2.Only = []string{"a", "b", "c"}

	opt2.assign(default2, file2, cli2)

	assert.Equal(t, "https://abc.com", opt2.Host, "should keep the Host")
	assert.Equal(t, []string{"a", "b", "c"}, opt2.Only, "should add Only to the option")

	opt3 := NewConfig()
	default3 := *NewConfig()
	file3 := Config{}
	cli3 := Config{}

	cli3.Host = "https://abc123.com"
	cli3.OnlyCLI = "x,y,z"

	opt3.assign(default3, file3, cli3)

	assert.Equal(t, "https://abc123.com", opt3.Host, "should keep the Host")
	assert.Equal(t, "x,y,z", opt3.OnlyCLI, "should add OnlyCLI to the option")
	assert.Equal(t, []string{"x", "y", "z"}, opt3.Only, "should add Only to the option - already parsed")
}

func TestConfigAssignOnlyIgnores(t *testing.T) {
	for _, v := range tableTestAssignIgnorePaths {
		opt := NewConfig()
		default1 := *NewConfig()
		file1 := Config{Ignore: v.inIgnoreFile}
		cli1 := Config{IgnoreCLI: v.inIgnoreCLI}

		opt.assign(default1, file1, cli1)

		assert.Equal(t, v.outIgnore, opt.Ignore, v.description)
	}
}

func TestReadIndexHTML(t *testing.T) {
	for _, v := range tableTestReadIndexHTML {
		m := myIndexFileMock{v.inInfo, v.inError}

		cfg := NewConfig()
		err := cfg.readIndexHTML(m)

		if (err != nil && v.outError != nil) && (err.Error() != v.outError.Error()) {
			assert.Fail(t, "should not fail now")
		}

		assert.Equal(t, v.outInfo, cfg.indexHTMLContent, v.description)
		assert.Equal(t, v.outError, err, v.description)
	}
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
		descriptionHTTP: "http - custom parsed value - not http2",
		outWSURL:        "ws://abc:9876/ws",
		descriptionWS:   "ws - custom parsed value - not secure",
	},
	{
		inHost:          "def",
		inPort:          ":1234",
		inSecure:        true,
		outHTTPURL:      "https://def:1234",
		descriptionHTTP: "http - custom parsed value - http2",
		outWSURL:        "wss://def:1234/ws",
		descriptionWS:   "ws - default parsed value - secure",
	},
}

var tableTestParseOnlyPaths = []struct {
	inOnly    []string
	inOnlyCLI string
	inRoot    string

	outOnly []string

	description string
}{
	{
		inOnly:    []string{},
		inOnlyCLI: "",
		inRoot:    "",

		outOnly:     []string{"."},
		description: "only default value - single dot",
	},
	{
		inOnly:    []string{},
		inOnlyCLI: "a",
		inRoot:    "",

		outOnly:     []string{"a"},
		description: "single onlyCLI value being parsed into Only",
	},
	{
		inOnly:    []string{},
		inOnlyCLI: "a,b,c",
		inRoot:    "",

		outOnly:     []string{"a", "b", "c"},
		description: "multiple onlyCLI value being parsed into Only",
	},
	{
		inOnly:    []string{"x", "y", "z"},
		inOnlyCLI: "a,b,c",
		inRoot:    "",

		outOnly:     []string{"x", "y", "z"},
		description: "Only value being left as it is",
	},
	{
		inOnly:    []string{"x", "y", "z"},
		inOnlyCLI: "",
		inRoot:    "base_root",

		outOnly:     []string{filepath.Join("base_root", "x"), filepath.Join("base_root", "y"), filepath.Join("base_root", "z")},
		description: "Only value should have the root as base",
	},
	{
		inOnly:    []string{"x", "y", "z"},
		inOnlyCLI: "1,2,3",
		inRoot:    "base_root2",

		outOnly:     []string{filepath.Join("base_root2", "x"), filepath.Join("base_root2", "y"), filepath.Join("base_root2", "z")},
		description: "Only value should have the root as base - it should ignore the values in onlyCLI",
	},
	{
		inOnly:    []string{},
		inOnlyCLI: "1,2,3",
		inRoot:    "base_root3",

		outOnly:     []string{filepath.Join("base_root3", "1"), filepath.Join("base_root3", "2"), filepath.Join("base_root3", "3")},
		description: "Only value should have the root as base - since Only is empty, it should parse and use the values in onlyCLI",
	},
	{
		inOnly:    []string{},
		inOnlyCLI: "",
		inRoot:    "base_root4",

		outOnly:     []string{"base_root4"},
		description: "Since both Only and OnlyCLI are empty, it should only use the root",
	},
	{
		inOnly:    []string{"a/b/**/*"},
		inOnlyCLI: "",
		inRoot:    "",

		outOnly:     []string{filepath.Join("a/b")},
		description: "single from the file - Removes * and removes duplication of * slashes",
	},
	{
		inOnly:    []string{"a/b/**/*", "x/**"},
		inOnlyCLI: "",
		inRoot:    "",

		outOnly:     []string{filepath.Join("a/b"), filepath.Join("x")},
		description: "multiple from the file - Removes * and removes duplication of * slashes",
	},
	{
		inOnly:    []string{"a/b/**/*", "x/**"},
		inOnlyCLI: "",
		inRoot:    "xabc",

		outOnly:     []string{filepath.Join("xabc/a/b"), filepath.Join("xabc/x")},
		description: "multiple from the file - with Root - Removes * and removes duplication of * slashes",
	},
	{
		inOnly:    []string{"/**/*"},
		inOnlyCLI: "",
		inRoot:    "",

		outOnly:     []string{"."},
		description: "multiple from the file - with Root - should leave the dot",
	},
	{
		inOnly:    []string{"/**/*"},
		inOnlyCLI: "",
		inRoot:    "abc",

		outOnly:     []string{"abc"},
		description: "multiple from the file - with Root - should leave the root",
	},
	{
		inOnly:    []string{},
		inOnlyCLI: "a/b/**/*",
		inRoot:    "",

		outOnly:     []string{filepath.Join("a/b")},
		description: "single from the CLI - Removes * and removes duplication of * slashes",
	},
	{
		inOnly:    []string{},
		inOnlyCLI: "a/b/**/*,x/**/*",
		inRoot:    "",

		outOnly:     []string{filepath.Join("a/b"), filepath.Join("x")},
		description: "multiple from the CLI - Removes * and removes duplication of * slashes",
	},
	{
		inOnly:    []string{},
		inOnlyCLI: "a/b/**/*,x/**/*",
		inRoot:    "x123",

		outOnly:     []string{filepath.Join("x123/a/b"), filepath.Join("x123/x")},
		description: "multiple from the CLI - with Root - Removes * and removes duplication of * slashes",
	},
	{
		inOnly:    []string{},
		inOnlyCLI: "**/*",
		inRoot:    "",

		outOnly:     []string{"."},
		description: "multiple from the CLI - should leave the dot only",
	},
	{
		inOnly:    []string{},
		inOnlyCLI: "**/*",
		inRoot:    "abc",

		outOnly:     []string{"abc"},
		description: "multiple from the CLI - should leave the root",
	},
}

var tableTestAssignIgnorePaths = []struct {
	inIgnoreCLI  string
	inIgnoreFile []string

	outIgnore []string

	description string
}{
	{
		inIgnoreCLI:  "",
		inIgnoreFile: []string{},
		outIgnore:    []string{},
		description:  "Both Ignore and IgnoreCLI are empty, so it should stay that way",
	},
	{
		inIgnoreCLI:  "a",
		inIgnoreFile: []string{},
		outIgnore:    []string{"a"},
		description:  "single - IgnoreFile is empty, so the cli should be the only path ignored",
	},
	{
		inIgnoreCLI:  "a,b,c",
		inIgnoreFile: []string{},
		outIgnore:    []string{"a", "b", "c"},
		description:  "multiple - IgnoreFile is empty, so the cli should be the only path ignored",
	},
	{
		inIgnoreCLI:  "",
		inIgnoreFile: []string{"a"},
		outIgnore:    []string{"a"},
		description:  "single - IgnoreCLI is empty, should use IgnoreFile instead",
	},
	{
		inIgnoreCLI:  "",
		inIgnoreFile: []string{"a", "b", "c"},
		outIgnore:    []string{"a", "b", "c"},
		description:  "multiple - IgnoreCLI is empty, should use IgnoreFile instead",
	},
	{
		inIgnoreCLI:  "x,y,z",
		inIgnoreFile: []string{"a", "b", "c"},
		outIgnore:    []string{"x", "y", "z"},
		description:  "multiple - IgnoreCLI should override the info in IgnoreFile",
	},
	{
		inIgnoreCLI:  "x,y,z",
		inIgnoreFile: []string{"a", "b", "c"},
		outIgnore:    []string{"x", "y", "z"},
		description:  "multiple - IgnoreCLI should override the info in IgnoreFile",
	},
	{
		inIgnoreCLI:  "x,y,z",
		inIgnoreFile: []string{"a", "b", "c"},
		outIgnore:    []string{"x", "y", "z"},
		description:  "multiple - IgnoreCLI should override the info in IgnoreFile",
	},
}

var tableTestParseIgnorePaths = []struct {
	inIgnoreCLI  string
	inIgnoreFile []string
	inRoot       string

	outIgnore []string

	description string
}{
	{
		inIgnoreCLI:  "",
		inIgnoreFile: []string{},
		inRoot:       "",
		outIgnore:    []string{},
		description:  "Both Ignore and IgnoreCLI are empty, so it should stay that way",
	},
	{
		inIgnoreCLI:  "a",
		inIgnoreFile: []string{},
		inRoot:       "",
		outIgnore:    []string{"a"},
		description:  "single - IgnoreFile is empty, so the cli should be the only path ignored",
	},
	{
		inIgnoreCLI:  "a,b,c",
		inIgnoreFile: []string{},
		inRoot:       "",
		outIgnore:    []string{"a", "b", "c"},
		description:  "multiple - IgnoreFile is empty, so the cli should be the only path ignored",
	},
	{
		inIgnoreCLI:  "",
		inIgnoreFile: []string{"a"},
		inRoot:       "",
		outIgnore:    []string{"a"},
		description:  "single - IgnoreCLI is empty, should use IgnoreFile instead",
	},
	{
		inIgnoreCLI:  "",
		inIgnoreFile: []string{"a", "b", "c"},
		inRoot:       "",
		outIgnore:    []string{"a", "b", "c"},
		description:  "multiple - IgnoreCLI is empty, should use IgnoreFile instead",
	},
	{
		inIgnoreCLI:  "x,y,z",
		inIgnoreFile: []string{"a", "b", "c"},
		inRoot:       "",
		outIgnore:    []string{"x", "y", "z"},
		description:  "multiple - IgnoreCLI should override the info in IgnoreFile",
	},
	{
		inIgnoreCLI:  "x,y,z",
		inIgnoreFile: []string{"a", "b", "c"},
		inRoot:       "",
		outIgnore:    []string{"x", "y", "z"},
		description:  "multiple - IgnoreCLI should override the info in IgnoreFile",
	},
	{
		inIgnoreCLI:  "x,y,z",
		inIgnoreFile: []string{"a", "b", "c"},
		inRoot:       "",
		outIgnore:    []string{"x", "y", "z"},
		description:  "multiple - IgnoreCLI should override the info in IgnoreFile",
	},
	{
		inIgnoreCLI:  "x,y,z",
		inIgnoreFile: []string{"a", "b", "c"},
		inRoot:       "zzz",
		outIgnore:    []string{filepath.Join("zzz", "x"), filepath.Join("zzz", "y"), filepath.Join("zzz", "z")},
		description:  "multiple - with root- IgnoreCLI should override the info in IgnoreFile",
	},
}

var tableTestParseIndexHTMLPathInfo = []struct {
	inRoot           string
	inPathIndex      string
	outIndexHTMLPath string
	description      string
}{
	{
		inRoot:           "",
		inPathIndex:      "",
		outIndexHTMLPath: "index.html",
		description:      "should have the index.html in the root of the app",
	},
	{
		inRoot:           "abc",
		inPathIndex:      "",
		outIndexHTMLPath: filepath.Join("", "abc", "index.html"),
		description:      "should have the index.html in the shallow folder - only root defined",
	},
	{
		inRoot:           "",
		inPathIndex:      "cde",
		outIndexHTMLPath: filepath.Join("", "cde", "index.html"),
		description:      "should have the index.html in the shallow folder - only pathIndex defined",
	},
	{
		inRoot:           "abc",
		inPathIndex:      "cde",
		outIndexHTMLPath: filepath.Join("", "abc", "cde", "index.html"),
		description:      "should have the index.html in deep folders - both root and pathIndex are defined",
	},
}

var tableTestReadIndexHTML = []struct {
	inInfo      string
	inError     error
	outInfo     []byte
	outError    error
	description string
}{
	{
		inInfo:      "abc",
		inError:     nil,
		outInfo:     []byte("abc"),
		outError:    nil,
		description: "should return the []byte correctly",
	},
	{
		inInfo:      "",
		inError:     errors.New("erro"),
		outInfo:     []byte(""),
		outError:    errors.New("erro"),
		description: "should return the error correctly",
	},
}
