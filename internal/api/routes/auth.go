package routes

import (
	"github.com/uptrace/bunrouter"

	"metalloyCore/internal/api/handler"
	"metalloyCore/internal/config"
	"metalloyCore/internal/database"
	"metalloyCore/internal/domain/auth"
	"metalloyCore/internal/domain/user"
	"metalloyCore/internal/security/jwt"
)

func AuthRoutes(cfg config.Setting) func(g *bunrouter.CompatGroup) {
	repository := user.InitRepository(cfg)
	userService := user.InitUserService(repository)
	jwtHandler := jwt.InitJWTHandler(cfg)
	redis := database.GetRedisClient(cfg)
	service := auth.InitAuthService(userService, jwtHandler, redis)
	controller := handler.InitAuthController(*service)
	return func(g *bunrouter.CompatGroup) {
		g.POST("/login", controller.LoginHandler)
		g.POST("/login-verify", controller.LoginVerifyHandler)
		g.POST("/register", controller.RegisterHandler)
		g.POST("/register-verify", controller.RegisterVerifyHandler)
		g.POST("/forgot-password", controller.ForgotPasswordHandler)
	}
}
