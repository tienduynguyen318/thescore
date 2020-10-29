package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"thescore/pkg/testutils"
)

var handler = func(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}

type fields struct {
	Method      string `json:"http.request.method"`
	Status      string `json:"http.response.status_code"`
	HTTPVersion string `json:"http.version"`
	Path        string `json:"url.full"`
	Message     string `json:"message"`
}

func TestRequestLogger(t *testing.T) {
	var buf bytes.Buffer
	l := testutils.SetupTestLogger(t, &buf)
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/some-path", nil)

	h := NewRequestLogger(l, http.HandlerFunc(handler))
	h.ServeHTTP(w, r)
	res := w.Result()

	if res.StatusCode != http.StatusNoContent {
		t.Errorf(
			"Request Logger handler not passing request down to concrete handler.\nExpected Status: %d\nGot: %d\n",
			http.StatusNoContent,
			res.StatusCode,
		)
	}

	expected := fields{
		Method:      "GET",
		Status:      "204 No Content",
		HTTPVersion: "HTTP/1.1",
		Path:        "/some-path",
	}
	var actual fields
	if err := json.Unmarshal(buf.Bytes(), &actual); err != nil {
		t.Fatalf("Failed to unmarshal structure log; %s", err)
	}
	pHost := `192\.0\.2\.1:1234 \- \-`
	pTimestamp := `\[\d{2}/[A-Z][a-z]{2}/\d{4}:\d{2}:\d{2}:\d{2} [-+]\d{4}\]`
	pHTTP := `GET /some-path HTTP/1\.1`
	pattern := fmt.Sprintf(
		`^%s %s "%s" 204 0$`,
		pHost,
		pTimestamp,
		pHTTP,
	)

	if m, err := regexp.MatchString(pattern, actual.Message); !m || err != nil {
		t.Errorf(
			"Message fields does not match.\nExpected Pattern:\n%s\nGot:\n%s\n",
			pattern,
			actual.Message,
		)
	}
	actual.Message = ""

	if actual != expected {
		t.Errorf(
			"Structure fields do not match expected fields.\nExpected:\n%+v\nGot:\n%+v\n",
			expected,
			actual,
		)
	}
}
