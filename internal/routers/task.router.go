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

	taskRoute := g.Group("/tasks")
	{
		taskRoute.POST("", handler.RegisterTasks)
		taskRoute.GET("", handler.GetTasks)
		taskRoute.GET("/:id", handler.GetTasksById)
		taskRoute.PUT("/:id", handler.EditTasks)
		taskRoute.DELETE("/:id", handler.DeleteTasks)
	}
}
