package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func Routrer(g *gin.Engine, db *sqlx.DB) {
	TaskRouter(g, db)
}
