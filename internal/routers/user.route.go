package routers

import (
	"github.com/funukonta/task_manager/internal/handlers"
	"github.com/funukonta/task_manager/internal/repositories"
	"github.com/funukonta/task_manager/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func UserRouter(g *gin.Engine, db *sqlx.DB) {
	repo := repositories.New_RepoUser(db)
	service := services.New_ServiceUser(repo)
	handler := handlers.New_HanlderUser(service)

	userRoute := g.Group("/users")
	{
		userRoute.POST("", handler.RegisterUser)
		userRoute.GET("", handler.GetUsers)
	}
}
