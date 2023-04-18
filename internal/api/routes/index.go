package routes

import "github.com/uptrace/bunrouter"

func V1Routes(g *bunrouter.CompatGroup) {
	g.WithGroup("", BaseRoutes)
	g.WithGroup("/user", UserRoutes)
}