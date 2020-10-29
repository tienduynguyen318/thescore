// A collection of utils that can be used by testing packages
package testutils

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"testing"
)

func VerifyGoldenResponse(t *testing.T, res *http.Response, f string) {
	expected, err := ioutil.ReadFile(f)
	if err != nil {
		t.Fatalf(`Unable to open "%s". Got: %s`, f, err)
	}
	expected = bytes.TrimSpace(expected)
	b, err := httputil.DumpResponse(res, true)
	if err != nil {
		t.Fatalf(`Unable to dump http response. %s`, err)
	}
	b = bytes.TrimSpace(
		bytes.ReplaceAll(b, []byte("\r\n"), []byte("\n")),
	)
	if !bytes.Equal(b, expected) {
		t.Errorf(
			"HTTP Response does not match Template: %s.\nExpected:\n%s\n\nGot:\n%s\n",
			f,
			expected,
			b,
		)
	}
}
