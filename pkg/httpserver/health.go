package httpserver

import (
	"net/http"
)

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		if s.healthHandlerOverride != nil {
			s.healthHandlerOverride(w, r)
			return
		}
		s.JSONResponse(w, r, map[string]string{"status": "ok"})
		return
	}
	http.NotFoundHandler().ServeHTTP(w, r)
}

func (s *Server) readyHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		if s.readyHandlerOverride != nil {
			s.readyHandlerOverride(w, r)
			return
		}
		s.JSONResponse(w, r, map[string]string{"status": "ok"})
		return
	}
	http.NotFoundHandler().ServeHTTP(w, r)
}
