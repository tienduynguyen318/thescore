package middleware

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"thescore/pkg/testutils"
)

var panicHandler = func(w http.ResponseWriter, r *http.Request) {
	panic(errors.New("Panic Error"))
}

var panicStringHandler = func(w http.ResponseWriter, r *http.Request) {
	panic("Panic String Error")
}

type StringerError struct{}

func (s StringerError) String() string {
	return "Panic Stringer Error"
}

var panicStringerHandler = func(w http.ResponseWriter, r *http.Request) {
	panic(StringerError{})
}

var panicUnknownHandler = func(w http.ResponseWriter, r *http.Request) {
	panic(struct{}{})
}

func TestRecoverMiddleware(t *testing.T) {
	cases := []struct {
		message string
		handler func(w http.ResponseWriter, r *http.Request)
	}{
		{message: "Panic Error", handler: panicHandler},
		{message: "Panic String Error", handler: panicStringHandler},
		{message: "Panic Stringer Error", handler: panicStringerHandler},
		{message: "An unknown error has occurred", handler: panicUnknownHandler},
	}
	for _, tc := range cases {
		t.Run(tc.message, func(t *testing.T) {
			var buf bytes.Buffer
			l := testutils.SetupTestLogger(t, &buf)

			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, "/some-path", nil)

			recoveryHandler := NewRecoverHandler(l, http.HandlerFunc(tc.handler))
			recoveryHandler.ServeHTTP(w, r)
			res := w.Result()

			testutils.VerifyGoldenResponse(t, res, "testdata/internal_server_error.golden")

			fields := []testutils.LogField{
				{Key: "log.level", Value: "error"},
				{Key: "message", Value: tc.message},
				{Key: "error.message", Value: tc.message},
			}
			testutils.ParseAndVerifyLogMessage(t, buf.Bytes(), fields)

		})
	}
}
