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
		g.GET("/:username", controller.UserHandler)
		// g.PUT("/:username", handler.UpdateUserHandler)
		// g.DELETE("/:username", handler.DeleteUserHandler)
	}
}
