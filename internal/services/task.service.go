package services

import (
	"fmt"
	"strings"

	"github.com/funukonta/task_manager/internal/models"
	"github.com/funukonta/task_manager/internal/repositories"
)

type Service_Task interface {
	RegisterTasks(*models.TasksModel) (*models.TasksModel, error)
	GetTasks() ([]models.TasksModel, error)
	GetTasksById(id int) (*models.TasksModel, error)
	EditTasks(data *models.TasksModel) error
	DeleteTasks(id int) error
}

type service_Task struct {
	repo repositories.Repo_Tasks
}

func New_ServiceTask(repo repositories.Repo_Tasks) Service_Task {
	return &service_Task{repo: repo}
}

func (s *service_Task) RegisterTasks(data *models.TasksModel) (*models.TasksModel, error) {
	return s.repo.RegisterTasks(data)
}
func (s *service_Task) GetTasks() ([]models.TasksModel, error) {
	result, err := s.repo.GetTasks()
	if err != nil {
		if strings.Contains(err.Error(), "no rows") {
			return nil, fmt.Errorf("tidak ada data")
		}
	}

	if len(result) == 0 {
		return nil, fmt.Errorf("tidak ada data")
	}

	return result, err
}
func (s *service_Task) GetTasksById(id int) (*models.TasksModel, error) {
	result, err := s.repo.GetTasksById(id)
	if err != nil {
		if strings.Contains(err.Error(), "no rows") {
			return nil, fmt.Errorf("tidak ada data")
		}
	}

	if result == nil {
		return nil, fmt.Errorf("tidak ada data")
	}

	return result, nil
}
func (s *service_Task) EditTasks(data *models.TasksModel) error {
	return nil
}
func (s *service_Task) DeleteTasks(id int) error {
	return nil
}