package httprest

import (
	"encoding/json"
	"net/http"
	"strings"
	"thescore/internal/domain"
	"thescore/pkg/httpserver"
	"thescore/pkg/logger"

	"github.com/julienschmidt/httprouter"
)

type Config struct {
	Service domain.Service
	Logger  logger.Logger
}

type Server struct {
	service domain.Service
	logger  logger.Logger
	router  *httprouter.Router
	*httpserver.Server
}

func New(config Config) (*Server, error) {
	var server Server
	s, err := httpserver.New(httpserver.Config{
		Logger: config.Logger,
		Addr:   "8080",
	})
	if err != nil {
		return nil, err
	}
	server.Server = s
	server.service = config.Service
	server.logger = config.Logger
	server.router = httprouter.New()
	server.setupRoutes()
	return &server, nil
}

// endpoint health and ready has been setup in httpserver package
func (s *Server) setupRoutes() {
	s.router.GET("/player", s.getPlayer)
	s.router.GET("/export", s.exportData)
	s.Server.HandleFunc("/", s.router.ServeHTTP)
}

func (s *Server) JSONErrorResponse(w http.ResponseWriter, r *http.Request, httpErr *httpError) {
	s.JSONResponseWithStatus(w, r, httpErr.Status, httpErr.Payload())
}

func (s *Server) MarshalPayload(w http.ResponseWriter, statusCode int, do json.Marshaler) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	rp, err := do.MarshalJSON()
	if err != nil {
		return err
	}
	if _, err := w.Write(rp); err != nil {
		return err
	}

	return nil
}

func (s *Server) parseNameParam(r *http.Request) []string {
	names := r.URL.Query()["name"]
	if len(names) > 0 {
		return strings.Split(names[0], ",")
	}
	return []string{}
}

func (s *Server) parseSortParam(r *http.Request) string {
	sortedBy := r.URL.Query().Get("sort")
	if len(sortedBy) > 0 {
		return sortedBy
	}
	return ""
}
