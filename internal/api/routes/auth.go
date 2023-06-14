package routes

import (
	"github.com/uptrace/bunrouter"

	"metalloyCore/internal/api/handler"
)

func AuthRoutes(controller *handler.AuthController) func(g *bunrouter.CompatGroup) {
	return func(g *bunrouter.CompatGroup) {
		g.POST("/login", controller.LoginHandler)
		g.POST("/login-verify", controller.LoginVerifyHandler)
		g.POST("/register", controller.RegisterHandler)
		g.POST("/register-verify", controller.RegisterVerifyHandler)
		g.POST("/reset-password", controller.ResetPasswordHandler)
		g.POST("/reset-password-verify", controller.ResetPasswordVerifyHandler)
		g.POST("/reset-password-final", controller.ResetPasswordFinalHandler)
	}
}
