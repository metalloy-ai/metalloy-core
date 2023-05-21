package routes

import (
	"github.com/uptrace/bunrouter"

	"metalloyCore/internal/api/handler"
	"metalloyCore/internal/config"
	"metalloyCore/internal/domain/user"
	"metalloyCore/internal/security/auth"
	"metalloyCore/internal/security/jwt"
)

func AuthRoutes(cfg config.Setting) func(g *bunrouter.CompatGroup) {
	repository := user.InitRepository(cfg)
	userService := user.InitUserService(repository)
	jwtHandler := jwt.InitJWTHandler(cfg)
	service := auth.InitAuthService(userService, jwtHandler)
	controller := handler.InitAuthController(*service)
	return func(g *bunrouter.CompatGroup) {
		g.POST("/login", controller.LoginHandler)
		g.POST("/register", controller.RegisterHandler)
		g.POST("/forgotPassword", controller.ForgotPasswordHandler)
	}
}
