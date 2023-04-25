package routes

import (
	"metalloyCore/internal/api/handler"
	"metalloyCore/internal/config"
	"metalloyCore/internal/domain/user"
	"metalloyCore/internal/security/auth"

	"github.com/uptrace/bunrouter"
)

func AuthRoutes(cfg config.Setting) func(g *bunrouter.CompatGroup) {
	repository := user.InitRepository(cfg)
	service := auth.InitAuthService(repository)
	controller := handler.InitAuthController(*service)
	return func(g *bunrouter.CompatGroup) {
		g.POST("/login", controller.LoginHandler)
		g.POST("/register", controller.RegisterHandler)
	}
}
