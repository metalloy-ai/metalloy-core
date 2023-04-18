package routes

import (
	"logiflowCore/internal/api/handler"

	"github.com/uptrace/bunrouter"
)

func BaseRoutes(group *bunrouter.Group) {
	group.GET("/", bunrouter.HTTPHandlerFunc(handler.BaseHandler))
	group.GET("/health", bunrouter.HTTPHandlerFunc(handler.HealthHandler))
}