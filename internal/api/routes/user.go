package routes

import (
	"github.com/uptrace/bunrouter"

	"metalloyCore/internal/api/handler"
	"metalloyCore/internal/config"
	"metalloyCore/internal/domain/user"
)

func UsersRoutes(cfg config.Setting) func(g *bunrouter.CompatGroup) {
	repository := user.InitRepository(cfg)
	controller := handler.InitUserController(repository)
	return func(g *bunrouter.CompatGroup) {
		g.GET("", controller.AllUserHandler)
		g.GET("/user", controller.EmptyParamHandler)
		g.WithGroup("/user/:username", UserRoutes(controller))
	}
}

func UserRoutes(controller *handler.UserController) func(g *bunrouter.CompatGroup) {
	return func(g *bunrouter.CompatGroup) {
		g.GET("", controller.UserHandler)
		g.PATCH("", controller.UpdateUserHandler)
		g.DELETE("", controller.DeleteUserHandler)
		g.WithGroup("/address", AddressRoutes(controller))
	}
}

func AddressRoutes(controller *handler.UserController) func(g *bunrouter.CompatGroup) {
	return func(g *bunrouter.CompatGroup) {
		g.GET("", controller.GetAddressHandler)
		g.PATCH("", controller.UpdateAddressHandler)
	}
}
