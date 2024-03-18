package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func UserRouter(g *gin.Engine, db *sqlx.DB) {

	userRoute := g.Group("/users")
	{
		userRoute.POST("")
	}
}
