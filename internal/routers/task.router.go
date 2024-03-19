package routers

import (
	"github.com/funukonta/task_manager/internal/handlers"
	"github.com/funukonta/task_manager/internal/repositories"
	"github.com/funukonta/task_manager/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func TaskRouter(g *gin.Engine, db *sqlx.DB) {
	repo := repositories.New_RepoTask(db)
	service := services.New_ServiceTask(repo)
	handler := handlers.New_HandlerTask(service)

	taskRoute := g.Group("/task")
	{
		taskRoute.POST("", handler.RegisterTasks)
	}
}
