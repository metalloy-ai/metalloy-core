package routes

import (
	"github.com/uptrace/bunrouter"

	"metalloyCore/internal/api/handler"
	"metalloyCore/internal/config"
	"metalloyCore/internal/domain/user"
	"metalloyCore/internal/security/auth"
)

func AuthRoutes(cfg config.Setting) func(g *bunrouter.CompatGroup) {
	repository := user.InitRepository(cfg)
	userService := user.InitUserService(repository)
	service := auth.InitAuthService(userService)
	controller := handler.InitAuthController(*service)
	return func(g *bunrouter.CompatGroup) {
		g.POST("/login", controller.LoginHandler)
		g.POST("/register", controller.RegisterHandler)
		g.POST("/forget-password", controller.ForgetPasswordHandler)
	}
}
