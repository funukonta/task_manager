package repositories

import "github.com/funukonta/task_manager/internal/models"

type Repo_Users interface {
	RegisterUser(*models.UserModel) (*models.UserModel, error)
	GetUsers() ([]models.UserModel, error)
	GetUserById(id int) (*models.UserModel, error)
	EditUser(data *models.UserModel) error
	DeleteUser(id int) error
}

type Repo_Tasks interface {
	RegisterTasks(*models.TasksModel) (*models.TasksModel, error)
	GetTasks() ([]models.TasksModel, error)
	GetTasksById(id int) (*models.TasksModel, error)
	// EditTasks(data *models.TasksModel) error
	// DeleteTasks(id int) error
}
