package routes

import (
	"logiflowCore/internal/api/handler"
	"logiflowCore/internal/config"

	"github.com/uptrace/bunrouter"
)

func BaseRoutes(cfg config.Setting) func(g *bunrouter.CompatGroup) {
	return func(g *bunrouter.CompatGroup) {
		g.GET("", handler.DefaultHandler)
		g.GET("/health", handler.HealthHandler(cfg))
	}
}
