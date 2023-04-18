package routes

import (
	"logiflowCore/internal/api/handler"

	"github.com/uptrace/bunrouter"
)

func BaseRoutes(g *bunrouter.CompatGroup) {
	g.GET("", handler.BaseHandler)
	g.GET("/health", handler.HealthHandler)
}