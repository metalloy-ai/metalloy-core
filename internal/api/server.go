package api

import (
	"logiflowCore/internal/config"
	"net/http"
	"strconv"

	"github.com/uptrace/bunrouter"
)

type Server struct {
	router 	   *bunrouter.Router
	host       string
	port   	   int
	apiVersion string
}

func InitServer() *Server {
	return &Server{
		router: bunrouter.New(),
		host: config.Host,
		port: config.Port,
		apiVersion: config.ApiVersion,
	}
}

func (s *Server) Run() {
	url := s.host + ":" + strconv.Itoa(s.port)
	http.ListenAndServe(url, s.router)
}

func (s *Server) LoadConfig() {
	
}