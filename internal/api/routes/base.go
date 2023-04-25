package routes

import (
	"metalloyCore/internal/api/handler"
	"metalloyCore/internal/config"

	"github.com/uptrace/bunrouter"
)

func BaseRoutes(cfg config.Setting) func(g *bunrouter.CompatGroup) {
	return func(g *bunrouter.CompatGroup) {
		g.GET("", handler.DefaultHandler)
		g.GET("/health", handler.HealthHandler(cfg))
	}
}
