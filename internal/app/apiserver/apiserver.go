package apiserver

import (
	"context"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// APIServer ...
type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	server *http.Server
}

// New ...
func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

// Start ...
func (s *APIServer) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}
	s.logger.Infof("Starting API Server at port %s", s.config.BindAddr)

	s.configureRouter()
	s.logger.Info("Routes configurete successfully")
	s.server = &http.Server{
		Addr:    s.config.BindAddr,
		Handler: s.router,
	}
	return s.server.ListenAndServe()
}

// Shutdown ...
func (s *APIServer) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}

// Logger ...
func (s *APIServer) Logger() *logrus.Logger {
	return s.logger
}

// configureLogger ...
func (s *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}
	s.logger.SetLevel(level)

	dir, err := filepath.Abs(filepath.Dir(s.config.LogFilePath))
	if err != nil {
		return err
	}
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.Mkdir(dir, os.ModePerm)
		if err != nil {
			return err
		}
	}

	file, err := os.OpenFile(s.config.LogFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	s.logger.SetOutput(io.MultiWriter(file, os.Stdout))
	return nil
}

// configureRouter ...
func (s *APIServer) configureRouter() {
	s.router.HandleFunc("/hello", s.handleHello())
}

// handleHello ...
func (s *APIServer) handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.logger.Infof("Method:%s URL:%s User-Agent:%s", r.Method, r.URL.String(), r.UserAgent())
		_, err := io.WriteString(w, "Hello, world!")
		if err != nil {
			s.logger.Errorf(
				"error occured while writing response to Method:%s URL:%s Error:%s",
				r.Method,
				r.URL.String(),
				err.Error(),
			)
		}
	}
}
