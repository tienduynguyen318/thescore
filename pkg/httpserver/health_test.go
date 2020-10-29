package httpserver

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"thescore/pkg/testutils/assert"
)

func TestHealthHandler(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/health", nil)
	resp := runServerTest(t, r)
	assert.HTTPStatus(t, resp, http.StatusOK)
	assert.HTTPHeader(t, resp, "Content-Type", "application/json")
	assert.HTTPBody(t, resp, `{"status":"ok"}`)
}

func TestHealthHandlerOverride(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/health", nil)
	resp := runServerTestWithOverrides(t, r)
	assert.HTTPStatus(t, resp, http.StatusOK)
	assert.HTTPHeader(t, resp, "Content-Type", "application/json")
	assert.HTTPBody(t, resp, `{"status":"override ok"}`)

}

func TestReadyHandler(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/ready", nil)
	resp := runServerTest(t, r)
	assert.HTTPStatus(t, resp, http.StatusOK)
	assert.HTTPHeader(t, resp, "Content-Type", "application/json")
	assert.HTTPBody(t, resp, `{"status":"ok"}`)
}

func TestReadyHandlerOverride(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/ready", nil)
	resp := runServerTestWithOverrides(t, r)
	assert.HTTPStatus(t, resp, http.StatusOK)
	assert.HTTPHeader(t, resp, "Content-Type", "application/json")
	assert.HTTPBody(t, resp, `{"status":"override ok"}`)
}
