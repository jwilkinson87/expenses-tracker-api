package http

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

type Router struct{}

func (r *Router) Setup(g *gin.Engine, db *sql.DB, httpHandlers []RouteRegistry) {
	for _, handler := range httpHandlers {
		handler.RegisterRoutes(g, db)
	}
}
