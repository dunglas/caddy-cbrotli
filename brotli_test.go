package caddycbrotli_test

import (
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/caddyserver/caddy/v2/caddytest"
	"github.com/google/brotli/go/cbrotli"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBrotli(t *testing.T) {
	const dummyDoc = `
<!doctype html>
<meta charset=utf-8>
<title>shortest html5</title>
Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.
Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.
Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.
Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.`

	tester := caddytest.NewTester(t)
	tester.InitServer(fmt.Sprintf(`
	{
		skip_install_trust
		admin localhost:2999
		http_port 9080
		https_port 9443
	}

	localhost:9080
	encode br
	header Content-Type text/html
	respond <<HTML
%s
HTML
	`, dummyDoc), "caddyfile")

	req, _ := http.NewRequest(http.MethodGet, "http://localhost:9080", nil)
	req.Header.Add("Accept-Encoding", "br")

	resp := tester.AssertResponseCode(req, http.StatusOK)
	defer resp.Body.Close()

	assert.Equal(t, "br", resp.Header.Get("Content-Encoding"))

	reader := cbrotli.NewReader(resp.Body)
	defer reader.Close()

	body, err := io.ReadAll(reader)
	require.NoError(t, err)
	assert.Equal(t, dummyDoc, string(body))
}
