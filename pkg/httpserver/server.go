package httpserver

import (
	"context"
	"net/http"

	"thescore/pkg/logger"
	"thescore/pkg/middleware"
)

type Config struct {
	Logger logger.Logger
	Addr   string
}

type Server struct {
	s                     http.Server
	logger                logger.Logger
	router                *http.ServeMux
	healthHandlerOverride http.HandlerFunc
	readyHandlerOverride  http.HandlerFunc
}

func New(config Config) (*Server, error) {
	var server Server
	server.logger = config.Logger
	server.setupRouter()
	server.s.Addr = ":8080"
	if config.Addr != "" {
		server.s.Addr = config.Addr
	}
	server.setupMiddleware()
	return &server, nil
}

func (s *Server) Handle(path string, handler http.Handler) {
	s.router.Handle(path, handler)
}

func (s *Server) HandleFunc(path string, handler func(http.ResponseWriter, *http.Request)) {
	s.router.HandleFunc(path, handler)
}

func (s *Server) setupRouter() {
	s.router = http.NewServeMux()
	s.router.HandleFunc("/health", s.healthHandler)
	s.router.HandleFunc("/ready", s.readyHandler)
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.s.Handler.ServeHTTP(w, r)
}
func (s *Server) setupMiddleware() {
	panicMiddleware := middleware.NewRecoverHandler(s.logger, s.router)
	logMiddleware := middleware.NewRequestLogger(s.logger, panicMiddleware)
	s.s.Handler = logMiddleware
}

func (s *Server) SetHealthHandler(h http.HandlerFunc) {
	s.healthHandlerOverride = h
}

func (s *Server) SetReadyHandler(h http.HandlerFunc) {
	s.readyHandlerOverride = h
}

func (s *Server) ListenAndServe() error {
	s.logger.Infof("Starting HTTP REST server")
	return s.s.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	s.logger.Infof("Shutting down HTTP REST server")
	return s.s.Shutdown(ctx)
}
