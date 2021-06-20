package apiserver

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
}

func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (s *APIServer) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	if err := s.configureRouter(); err != nil {
		return err
	}

	s.logger.Info("Starting API server, bind address: '", s.config.BindAddr, "'")

	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)

	return nil
}

func (s *APIServer) configureRouter() error {
	s.router.HandleFunc("/hello", s.handleHello())

	return nil
}

func (s *APIServer) handleHello() http.HandlerFunc {
	//

	return func(w http.ResponseWriter, r *http.Request) {
		_, _ = io.WriteString(w, "hello")
	}
}
