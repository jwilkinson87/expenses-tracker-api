package http

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

type RouteRegistry interface {
	RegisterRoutes(g *gin.Engine, db *sql.DB)
}
