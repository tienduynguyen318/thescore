package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
	"runtime/debug"

	"thescore/pkg/logger"
)

type HTTPError struct {
	Title  string `json:"title"`
	Status string `json:"status"`
	Detail string `json:"detail"`
}

type ErrorResponse struct {
	Errors []HTTPError `json:"errors"`
}

type RecoverHandler struct {
	l logger.Logger
	h http.Handler
}

// NewREcoveryHandler provides middleware that catches a "panic" call, logs the error
// and returns a 500 generic server response without causing the application to fail.
// Takes in a logger.Logger instance and a http.Handler that is wrapped.
//
// Idealing this would be the top level piece of middleware
func NewRecoverHandler(l logger.Logger, h http.Handler) *RecoverHandler {
	return &RecoverHandler{l: l, h: h}
}

func (rh *RecoverHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			var message string
			switch e := r.(type) {
			case error:
				message = e.Error()
			case fmt.Stringer:
				message = e.String()
			case string:
				message = e
			default:
				message = "An unknown error has occurred"
			}
			stacktrace := debug.Stack()
			rh.l.WithFields(
				logger.String("error.message", message),
				logger.String("error.stack_trace", string(stacktrace)),
			).Errorf(message)

			payload := ErrorResponse{
				Errors: []HTTPError{
					{
						Title:  "Internal Server Error",
						Status: "500",
						Detail: "An Internal Server Error has occurred.",
					},
				},
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			// nolint:errorcheck
			json.NewEncoder(w).Encode(payload)
		}
	}()
	rh.h.ServeHTTP(w, r)
}
