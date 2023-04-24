package api

import (
	"log"
	"net/http"
	"strconv"

	"logiflowCore/internal/api/routes"
	"logiflowCore/internal/config"
	"logiflowCore/internal/security/middleware"

	"github.com/uptrace/bunrouter"
	"github.com/uptrace/bunrouter/extra/reqlog"
)

type Server struct {
	router *bunrouter.CompatRouter
	config config.Setting
}

func InitServer(config config.Setting) *Server {
	options := []bunrouter.Option{}
	if config.Env == "dev" {
		options = append(options, bunrouter.Use(reqlog.NewMiddleware()))
	}
	
	newRouter := bunrouter.New(options...).Compat()
	return &Server{
		router: newRouter,
		config: config,
	}
}

func (s *Server) Run() {
	url := s.config.Host + ":" + strconv.Itoa(s.config.Port)
	log.Println("Server is running on " + url)
	log.Println(http.ListenAndServe(url, s.router))
}

func (s *Server) LoadServerConfig() {
	cors := middleware.CorsMiddleware
	s.router.
		Use(cors).
		Compat().
	WithGroup("/api/v1", routes.V1Routes(s.config))
}
