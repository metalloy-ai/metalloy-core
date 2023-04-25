package routes

import (
	"metalloyCore/internal/config"

	"github.com/uptrace/bunrouter"
)

func V1Routes(cfg config.Setting) func(g *bunrouter.CompatGroup) {
	return func(g *bunrouter.CompatGroup) {
		g.WithGroup("", BaseRoutes(cfg))
		g.WithGroup("/user", UserRoutes(cfg))
		g.WithGroup("/auth", AuthRoutes(cfg))
	}
}
