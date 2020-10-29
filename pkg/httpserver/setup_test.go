package httpserver

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"thescore/pkg/testutils"
)

func runServerTest(t *testing.T, r *http.Request) *http.Response {
	w := httptest.NewRecorder()
	logger := testutils.SetupTestLogger(t, ioutil.Discard)
	config := Config{
		Logger: logger,
	}
	server, err := New(config)
	if err != nil {
		t.Fatalf("Failed setting up HTTP Server: %v", err)
	}

	server.ServeHTTP(w, r)
	return w.Result()
}

func runServerTestWithOverrides(t *testing.T, r *http.Request) *http.Response {
	w := httptest.NewRecorder()
	logger := testutils.SetupTestLogger(t, ioutil.Discard)
	config := Config{
		Logger: logger,
	}
	server, err := New(config)
	if err != nil {
		t.Fatalf("Failed setting up HTTP Server: %v", err)
	}

	server.SetHealthHandler(server.healthOverride)
	server.SetReadyHandler(server.readyOverride)

	server.ServeHTTP(w, r)
	return w.Result()
}

func (s *Server) healthOverride(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		s.JSONResponse(w, r, map[string]string{"status": "override ok"})
		return
	}
	http.NotFoundHandler().ServeHTTP(w, r)
}

func (s *Server) readyOverride(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		s.JSONResponse(w, r, map[string]string{"status": "override ok"})
		return
	}
	http.NotFoundHandler().ServeHTTP(w, r)
}
