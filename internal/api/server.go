package api

import (
	"log"
	"net/http"
	"strconv"

	"logiflowCore/internal/config"
	"logiflowCore/internal/api/routes"

	"github.com/uptrace/bunrouter"
)

type Server struct {
	router 	   *bunrouter.CompatRouter
	host       string
	port   	   int
	apiVersion string
}

func InitServer() *Server {
	return &Server{
		router:   	bunrouter.New().Compat(),
		host:   	config.Host,
		port:   	config.Port,
		apiVersion: config.ApiVersion,
	}
}

func (s *Server) Run() {
	url := s.host + ":" + strconv.Itoa(s.port)
	log.Println("Server is running on " + url)
	log.Println(http.ListenAndServe(url, s.router))
}

func (s *Server) LoadConfig() {
	s.router.WithGroup("/api/v1", routes.V1Routes)
	
}