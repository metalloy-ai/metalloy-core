package routes

import (
	"logiflowCore/internal/api/handler"
	"logiflowCore/internal/config"
	"logiflowCore/internal/domain/user"

	"github.com/uptrace/bunrouter"
)

func UserRoutes(cfg config.Setting) func(g *bunrouter.CompatGroup) {
	repository := user.InitRepository(cfg)
	controller := handler.InitUserController(repository)
	return func(g *bunrouter.CompatGroup) {
		g.GET("", controller.AllUserHandler)
		g.GET("/:username", controller.SingleUserHandler)
		// g.POST("/register", handler.RegisterHandler)
		// g.POST("/login", handler.LoginHandler)
		// g.POST("/logout", handler.LogoutHandler)
		// g.PUT("/:username", handler.UpdateUserHandler)
		// g.DELETE("/:username", handler.DeleteUserHandler)
	}
}
