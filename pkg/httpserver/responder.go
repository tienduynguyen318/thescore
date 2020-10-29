package httpserver

import (
	"encoding/json"
	"net/http"
)

func (s *Server) JSONResponseWithStatus(w http.ResponseWriter, r *http.Request, status int, data interface{}) {
	body, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		s.logger.Errorf("JSON marshal failed")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(status)
	w.Write(body)

}

func (s *Server) JSONResponse(w http.ResponseWriter, r *http.Request, data interface{}) {
	s.JSONResponseWithStatus(w, r, http.StatusOK, data)
}
