package api

import (
	"log"
	"net/http"
	"strconv"

	"github.com/uptrace/bunrouter"
	"github.com/uptrace/bunrouter/extra/reqlog"

	"metalloyCore/internal/api/routes"
	"metalloyCore/internal/config"
	"metalloyCore/internal/security/middleware"
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
	s.LoadRoutes()
}

func (s *Server) LoadRoutes() {
	v1Group := s.router.NewGroup("/api/v1")
	v1Group.WithMiddleware(middleware.CorsMiddleware)

	v1Group.WithGroup("", routes.BaseRoutes(s.config))
	v1Group.WithGroup("/auth", routes.AuthRoutes(s.config))

	v1Group.WithMiddleware(middleware.Authorization(s.config)).
		WithGroup("/users", routes.UsersRoutes(s.config))
}
