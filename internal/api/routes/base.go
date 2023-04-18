package routes

import (
	"logiflowCore/internal/api/handler"

	"github.com/uptrace/bunrouter"
)

func BaseRoutes(group *bunrouter.CompatGroup) {
	group.GET("/", handler.BaseHandler)
	group.GET("/health", handler.HealthHandler)
}