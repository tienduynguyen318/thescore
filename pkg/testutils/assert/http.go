package assert

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"thescore/pkg/testutils"
)

var VerifyGoldenResponse = testutils.VerifyGoldenResponse

func HTTPStatus(t *testing.T, resp *http.Response, expected int) {
	t.Helper()
	if resp.StatusCode != expected {
		t.Errorf("resp.StatusCode != \"%d\". Got: \"%d\".", expected, resp.StatusCode)
	}
}

func HTTPHeader(t *testing.T, r *http.Response, name, value string) {
	t.Helper()
	hvalue := r.Header.Get(name)
	if hvalue != value {
		t.Errorf(`Header[%s] != "%s". Got: "%s"`, name, value, hvalue)
	}
}

func HTTPBody(t *testing.T, r *http.Response, expected string) {
	t.Helper()
	respBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON: %v", err)
	}
	defer r.Body.Close()
	if string(respBody) != expected {
		t.Errorf("Response Body does not match.\nExpected:\n%s\n\nGot:\n%s\n", expected, respBody)
	}
}

func JSONGoldenResponse(t *testing.T, fname string, resp *http.Response) {
	t.Helper()
	raw, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Unable to read HTTP Response body: %v", err)
	}
	JSONGolden(t, fname, raw)
}

func JSONGolden(t *testing.T, fname string, raw []byte) {
	t.Helper()
	var actual bytes.Buffer
	expected, err := ioutil.ReadFile(fname)
	if err != nil {
		t.Fatalf(`Unable to open "%s" golden file. Got: %v`, fname, err)
	}
	expected = bytes.TrimSpace(expected)
	if err := json.Indent(&actual, raw, "", "  "); err != nil {
		t.Fatalf("HTTP Response Body is not in json format: %v", err)
	}
	if !bytes.Equal(actual.Bytes(), expected) {
		t.Errorf(
			"JSON golden file (%s) does not match.\nExpected:\n%s\n\nGot:\n%s\n",
			fname,
			expected,
			actual.String(),
		)
	}

}
