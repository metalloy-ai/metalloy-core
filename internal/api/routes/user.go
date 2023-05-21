package routes

import (
	"metalloyCore/internal/api/handler"
	"metalloyCore/internal/config"
	"metalloyCore/internal/domain/user"
	"metalloyCore/internal/security/middleware"

	"github.com/uptrace/bunrouter"
)

func UsersRoutes(cfg config.Setting) func(g *bunrouter.CompatGroup) {
	repository := user.InitRepository(cfg)
	service := user.InitUserService(repository)
	controller := handler.InitUserController(service)
	return func(g *bunrouter.CompatGroup) {
		g.GET("", middleware.Admin(controller.AllUserHandler))
		g.GET("/user", controller.EmptyParamHandler)
		g.WithGroup("/user/:username", UserRoutes(controller))
	}
}

func UserRoutes(controller *handler.UserController) func(g *bunrouter.CompatGroup) {
	return func(g *bunrouter.CompatGroup) {
		g.GET("", controller.UserHandler)
		g.PATCH("", controller.UpdateUserHandler)
		g.DELETE("", middleware.Admin(controller.DeleteUserHandler))
		g.WithGroup("/address", AddressRoutes(controller))
	}
}

func AddressRoutes(controller *handler.UserController) func(g *bunrouter.CompatGroup) {
	return func(g *bunrouter.CompatGroup) {
		g.GET("", controller.GetAddressHandler)
		g.PATCH("", controller.UpdateAddressHandler)
	}
}
