package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func TaskRouter(g *gin.Engine, db *sqlx.DB) {
	taskRoute := g.Group("/task")
	{
		taskRoute.POST("")
	}
}
