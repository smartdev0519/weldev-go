// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package httpproxy_test

import (
	"bytes"
	"errors"
	"fmt"
	"net/url"
	"os"
	"strings"
	"testing"

	"golang_org/x/net/http/httpproxy"
)

type proxyForURLTest struct {
	cfg     httpproxy.Config
	req     string // URL to fetch; blank means "http://example.com"
	want    string
	wanterr error
}

func (t proxyForURLTest) String() string {
	var buf bytes.Buffer
	space := func() {
		if buf.Len() > 0 {
			buf.WriteByte(' ')
		}
	}
	if t.cfg.HTTPProxy != "" {
		fmt.Fprintf(&buf, "http_proxy=%q", t.cfg.HTTPProxy)
	}
	if t.cfg.HTTPSProxy != "" {
		space()
		fmt.Fprintf(&buf, "https_proxy=%q", t.cfg.HTTPSProxy)
	}
	if t.cfg.NoProxy != "" {
		space()
		fmt.Fprintf(&buf, "no_proxy=%q", t.cfg.NoProxy)
	}
	req := "http://example.com"
	if t.req != "" {
		req = t.req
	}
	space()
	fmt.Fprintf(&buf, "req=%q", req)
	return strings.TrimSpace(buf.String())
}

var proxyForURLTests = []proxyForURLTest{{
	cfg: httpproxy.Config{
		HTTPProxy: "127.0.0.1:8080",
	},
	want: "http://127.0.0.1:8080",
}, {
	cfg: httpproxy.Config{
		HTTPProxy: "cache.corp.example.com:1234",
	},
	want: "http://cache.corp.example.com:1234",
}, {
	cfg: httpproxy.Config{
		HTTPProxy: "cache.corp.example.com",
	},
	want: "http://cache.corp.example.com",
}, {
	cfg: httpproxy.Config{
		HTTPProxy: "https://cache.corp.example.com",
	},
	want: "https://cache.corp.example.com",
}, {
	cfg: httpproxy.Config{
		HTTPProxy: "http://127.0.0.1:8080",
	},
	want: "http://127.0.0.1:8080",
}, {
	cfg: httpproxy.Config{
		HTTPProxy: "https://127.0.0.1:8080",
	},
	want: "https://127.0.0.1:8080",
}, {
	cfg: httpproxy.Config{
		HTTPProxy: "socks5://127.0.0.1",
	},
	want: "socks5://127.0.0.1",
}, {
	// Don't use secure for http
	cfg: httpproxy.Config{
		HTTPProxy:  "http.proxy.tld",
		HTTPSProxy: "secure.proxy.tld",
	},
	req:  "http://insecure.tld/",
	want: "http://http.proxy.tld",
}, {
	// Use secure for https.
	cfg: httpproxy.Config{
		HTTPProxy:  "http.proxy.tld",
		HTTPSProxy: "secure.proxy.tld",
	},
	req:  "https://secure.tld/",
	want: "http://secure.proxy.tld",
}, {
	cfg: httpproxy.Config{
		HTTPProxy:  "http.proxy.tld",
		HTTPSProxy: "https://secure.proxy.tld",
	},
	req:  "https://secure.tld/",
	want: "https://secure.proxy.tld",
}, {
	// Issue 16405: don't use HTTP_PROXY in a CGI environment,
	// where HTTP_PROXY can be attacker-controlled.
	cfg: httpproxy.Config{
		HTTPProxy: "http://10.1.2.3:8080",
		CGI:       true,
	},
	want:    "<nil>",
	wanterr: errors.New("refusing to use HTTP_PROXY value in CGI environment; see golang.org/s/cgihttpproxy"),
}, {
	// HTTPS proxy is still used even in CGI environment.
	// (perhaps dubious but it's the historical behaviour).
	cfg: httpproxy.Config{
		HTTPSProxy: "https://secure.proxy.tld",
		CGI:        true,
	},
	req:  "https://secure.tld/",
	want: "https://secure.proxy.tld",
}, {
	want: "<nil>",
}, {
	cfg: httpproxy.Config{
		NoProxy:   "example.com",
		HTTPProxy: "proxy",
	},
	req:  "http://example.com/",
	want: "<nil>",
}, {
	cfg: httpproxy.Config{
		NoProxy:   ".example.com",
		HTTPProxy: "proxy",
	},
	req:  "http://example.com/",
	want: "<nil>",
}, {
	cfg: httpproxy.Config{
		NoProxy:   "ample.com",
		HTTPProxy: "proxy",
	},
	req:  "http://example.com/",
	want: "http://proxy",
}, {
	cfg: httpproxy.Config{
		NoProxy:   "example.com",
		HTTPProxy: "proxy",
	},
	req:  "http://foo.example.com/",
	want: "<nil>",
}, {
	cfg: httpproxy.Config{
		NoProxy:   ".foo.com",
		HTTPProxy: "proxy",
	},
	req:  "http://example.com/",
	want: "http://proxy",
}}

func testProxyForURL(t *testing.T, tt proxyForURLTest) {
	t.Helper()
	reqURLStr := tt.req
	if reqURLStr == "" {
		reqURLStr = "http://example.com"
	}
	reqURL, err := url.Parse(reqURLStr)
	if err != nil {
		t.Errorf("invalid URL %q", reqURLStr)
		return
	}
	cfg := tt.cfg
	proxyForURL := cfg.ProxyFunc()
	url, err := proxyForURL(reqURL)
	if g, e := fmt.Sprintf("%v", err), fmt.Sprintf("%v", tt.wanterr); g != e {
		t.Errorf("%v: got error = %q, want %q", tt, g, e)
		return
	}
	if got := fmt.Sprintf("%s", url); got != tt.want {
		t.Errorf("%v: got URL = %q, want %q", tt, url, tt.want)
	}

	// Check that changing the Config doesn't change the results
	// of the functuon.
	cfg = httpproxy.Config{}
	url, err = proxyForURL(reqURL)
	if g, e := fmt.Sprintf("%v", err), fmt.Sprintf("%v", tt.wanterr); g != e {
		t.Errorf("(after mutating config) %v: got error = %q, want %q", tt, g, e)
		return
	}
	if got := fmt.Sprintf("%s", url); got != tt.want {
		t.Errorf("(after mutating config) %v: got URL = %q, want %q", tt, url, tt.want)
	}
}

func TestProxyForURL(t *testing.T) {
	for _, tt := range proxyForURLTests {
		testProxyForURL(t, tt)
	}
}

func TestFromEnvironment(t *testing.T) {
	os.Setenv("HTTP_PROXY", "httpproxy")
	os.Setenv("HTTPS_PROXY", "httpsproxy")
	os.Setenv("NO_PROXY", "noproxy")
	os.Setenv("REQUEST_METHOD", "")
	got := httpproxy.FromEnvironment()
	want := httpproxy.Config{
		HTTPProxy:  "httpproxy",
		HTTPSProxy: "httpsproxy",
		NoProxy:    "noproxy",
	}
	if *got != want {
		t.Errorf("unexpected proxy config, got %#v want %#v", got, want)
	}
}

func TestFromEnvironmentWithRequestMethod(t *testing.T) {
	os.Setenv("HTTP_PROXY", "httpproxy")
	os.Setenv("HTTPS_PROXY", "httpsproxy")
	os.Setenv("NO_PROXY", "noproxy")
	os.Setenv("REQUEST_METHOD", "PUT")
	got := httpproxy.FromEnvironment()
	want := httpproxy.Config{
		HTTPProxy:  "httpproxy",
		HTTPSProxy: "httpsproxy",
		NoProxy:    "noproxy",
		CGI:        true,
	}
	if *got != want {
		t.Errorf("unexpected proxy config, got %#v want %#v", got, want)
	}
}

func TestFromEnvironmentLowerCase(t *testing.T) {
	os.Setenv("http_proxy", "httpproxy")
	os.Setenv("https_proxy", "httpsproxy")
	os.Setenv("no_proxy", "noproxy")
	os.Setenv("REQUEST_METHOD", "")
	got := httpproxy.FromEnvironment()
	want := httpproxy.Config{
		HTTPProxy:  "httpproxy",
		HTTPSProxy: "httpsproxy",
		NoProxy:    "noproxy",
	}
	if *got != want {
		t.Errorf("unexpected proxy config, got %#v want %#v", got, want)
	}
}

var UseProxyTests = []struct {
	host  string
	match bool
}{
	// Never proxy localhost:
	{"localhost", false},
	{"127.0.0.1", false},
	{"127.0.0.2", false},
	{"[::1]", false},
	{"[::2]", true}, // not a loopback address

	{"barbaz.net", false},     // match as .barbaz.net
	{"foobar.com", false},     // have a port but match
	{"foofoobar.com", true},   // not match as a part of foobar.com
	{"baz.com", true},         // not match as a part of barbaz.com
	{"localhost.net", true},   // not match as suffix of address
	{"local.localhost", true}, // not match as prefix as address
	{"barbarbaz.net", true},   // not match because NO_PROXY have a '.'
	{"www.foobar.com", false}, // match because NO_PROXY includes "foobar.com"
}

func TestUseProxy(t *testing.T) {
	cfg := &httpproxy.Config{
		NoProxy: "foobar.com, .barbaz.net",
	}
	for _, test := range UseProxyTests {
		if httpproxy.ExportUseProxy(cfg, test.host+":80") != test.match {
			t.Errorf("useProxy(%v) = %v, want %v", test.host, !test.match, test.match)
		}
	}
}

func TestInvalidNoProxy(t *testing.T) {
	cfg := &httpproxy.Config{
		NoProxy: ":1",
	}
	ok := httpproxy.ExportUseProxy(cfg, "example.com:80") // should not panic
	if !ok {
		t.Errorf("useProxy unexpected return; got false; want true")
	}
}
