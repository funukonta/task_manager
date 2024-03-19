package repositories

import "github.com/funukonta/task_manager/internal/models"

type Repo_Users interface {
	RegisterUser(*models.UserModel) (*models.UserModel, error)
	GetUsers() ([]models.UserModel, error)
	GetUserById(id int) (*models.UserModel, error)
	EditUser(id int)
	DeleteUser(id int)
}

type Repo_Tasks interface {
	AddTask()
	GetTasks()
	GetTaskById(id int)
	EditTask(id int)
	DeleteTask(id int)
}
