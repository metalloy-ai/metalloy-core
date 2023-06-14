package routes

import (
	"github.com/uptrace/bunrouter"

	"metalloyCore/internal/api/handler"
	"metalloyCore/internal/security/middleware"
)

func UsersRoutes(controller *handler.UserController) func(g *bunrouter.CompatGroup) {
	return func(g *bunrouter.CompatGroup) {
		g.GET("", middleware.Admin(controller.AllUserHandler))
		g.GET("/user", handler.EmptyParamHandler)
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
