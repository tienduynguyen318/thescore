package middleware

import (
	"fmt"
	"net/http"
	"time"

	"thescore/pkg/logger"
)

// NewRequestLogger returns a HTTP middleware that log the request using ECS (https://www.elastic.co/guide/en/ecs/current/index.html)
// derived fields. The message format is in the Common Log Format (https://en.wikipedia.org/wiki/Common_Log_Format)
//
// Also logs server panics including stacktracing.
//
// This should be the top middleware that is applied.
func NewRequestLogger(l logger.Logger, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		respWriter := &augmentedResponseWriter{
			ResponseWriter: w,
		}

		now := time.Now()
		h.ServeHTTP(respWriter, r)
		message := fmt.Sprintf(
			`%s - - [%s] "%s %s %s" %d %d`,
			r.RemoteAddr,
			now.Format("02/Jan/2006:15:04:05 -0700"),
			r.Method,
			r.URL.Path,
			r.Proto,
			respWriter.StatusCode,
			respWriter.Length,
		)
		reqLogger := l.WithFields(
			logger.String("http.request.method", r.Method),
			logger.String("http.version", r.Proto),
			logger.String("url.full", r.URL.Path),
			logger.String("http.response.status_code", fullStatusString(respWriter.StatusCode)),
		)
		reqLogger.Infof(message)
	})

}

type augmentedResponseWriter struct {
	http.ResponseWriter
	StatusCode int
	Length     int
}

func (w *augmentedResponseWriter) WriteHeader(statusCode int) {
	w.StatusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}
func (w *augmentedResponseWriter) Write(b []byte) (n int, err error) {
	n, err = w.ResponseWriter.Write(b)
	w.Length += n
	return
}

func fullStatusString(s int) string {
	return fmt.Sprintf("%d %s", s, http.StatusText(s))
}
