package goliv

import "testing"

var optDefaut = &Options{
	NoBrowser:   false,
	Host:        "127.0.0.1",
	Secure:      false,
	Port:        ":1308",
	PathIndex:   "",
	Quiet:       false,
	Proxy:       false,
	ProxyTarget: "",
	ProxyWhen:   "",
	Ignore:      "",
	Only:        ".",
}

func BenchmarkNewOptionsEmptyParams(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewOptions()
	}
}

func TestNewOptionsWithEmptyParamsShouldReturnDefault(t *testing.T) {
	o := NewOptions()

	if o.NoBrowser != optDefaut.NoBrowser {
		t.Errorf("Expected NoBrowser to be %v, but got %v", optDefaut.NoBrowser, o.NoBrowser)
	}

	if o.Host != optDefaut.Host {
		t.Errorf("Expected Host to be %s, but got %s", optDefaut.Host, o.Host)
	}

	if o.Secure != optDefaut.Secure {
		t.Errorf("Expected Secure to be %v, but got %v", optDefaut.Secure, o.Secure)
	}

	if o.Port != optDefaut.Port {
		t.Errorf("Expected Port to be %s, but got %s", optDefaut.Port, o.Port)
	}

	if o.PathIndex != optDefaut.PathIndex {
		t.Errorf("Expected PathIndex to be %s, but got %s", optDefaut.PathIndex, o.PathIndex)
	}

	if o.Quiet != optDefaut.Quiet {
		t.Errorf("Expected Quiet to be %v, but got %v", optDefaut.Quiet, o.Quiet)
	}

	if o.Proxy != optDefaut.Proxy {
		t.Errorf("Expected Proxy to be %v, but got %v", optDefaut.Proxy, o.Proxy)
	}

	if o.ProxyTarget != optDefaut.ProxyTarget {
		t.Errorf("Expected ProxyTarget to be %v, but got %v", optDefaut.ProxyTarget, o.ProxyTarget)
	}

	if o.ProxyWhen != optDefaut.ProxyWhen {
		t.Errorf("Expected ProxyWhen to be %s, but got %s", optDefaut.ProxyWhen, o.ProxyWhen)
	}

	if o.Ignore != optDefaut.Ignore {
		t.Errorf("Expected Ignore to be %s, but got %s", optDefaut.Ignore, o.Ignore)
	}

	if o.Only != optDefaut.Only {
		t.Errorf("Expected Only to be %s, but got %s", optDefaut.Only, o.Only)
	}
}
