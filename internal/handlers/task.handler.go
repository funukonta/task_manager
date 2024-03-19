package handlers

import (
	"net/http"

	"github.com/funukonta/task_manager/internal/models"
	"github.com/funukonta/task_manager/internal/services"
	"github.com/funukonta/task_manager/pkg"
	"github.com/gin-gonic/gin"
)

type handler_Task struct {
	service services.Service_Task
}

func New_HandlerTask(service services.Service_Task) *handler_Task {
	return &handler_Task{service: service}
}

func (h *handler_Task) RegisterTasks(c *gin.Context) {
	tasks := models.TasksModel{}
	if err := c.ShouldBind(&tasks); err != nil {
		pkg.Responses(http.StatusBadRequest, &pkg.BodJson{Message: err.Error()}).Send(c)
		return
	}

	result, err := h.service.RegisterTasks(&tasks)
	if err != nil {
		pkg.Responses(http.StatusBadRequest, &pkg.BodJson{Message: err.Error()}).Send(c)
		return
	}

	pkg.Responses(http.StatusOK, &pkg.BodJson{Data: result}).Send(c)
}

func (h *handler_Task) GetTasks(c *gin.Context) {
	result, err := h.service.GetTasks()
	if err != nil {
		pkg.Responses(http.StatusBadRequest, &pkg.BodJson{Message: err.Error()}).Send(c)
		return
	}

	pkg.Responses(http.StatusOK, &pkg.BodJson{Data: result}).Send(c)
}

func (h *handler_Task) GetTasksById(c *gin.Context) {
	task := &models.TasksModel{}
	if err := c.ShouldBind(task); err != nil {
		pkg.Responses(http.StatusBadRequest, &pkg.BodJson{Message: err.Error()}).Send(c)
		return
	}

	result, err := h.service.GetTasksById(task.ID)
	if err != nil {
		pkg.Responses(http.StatusBadRequest, &pkg.BodJson{Message: err.Error()}).Send(c)
		return
	}

	pkg.Responses(http.StatusOK, &pkg.BodJson{Data: result}).Send(c)
}

func (h *handler_Task) EditTasks(c *gin.Context) {
	task := &models.TasksModel{}
	if err := c.ShouldBind(task); err != nil {
		pkg.Responses(http.StatusBadRequest, &pkg.BodJson{Message: err.Error()}).Send(c)
		return
	}
	id := c.Param("id")

	if err := h.service.EditTasks(id, task); err != nil {
		pkg.Responses(http.StatusBadRequest, &pkg.BodJson{Message: err.Error()}).Send(c)
		return
	}

	pkg.Responses(http.StatusOK, &pkg.BodJson{Message: "Berhasil edit tasks"}).Send(c)
}
