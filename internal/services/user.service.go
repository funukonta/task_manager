package services

import (
	"fmt"

	"github.com/funukonta/task_manager/internal/models"
	"github.com/funukonta/task_manager/internal/repositories"
)

type Service_User interface {
	RegisterUser(*models.UserModel) (*models.UserModel, error)
	GetUsers() ([]models.UserModel, error)
	GetUserById(id int) (*models.UserModel, error)
	EditUser(data *models.UserModel) error
	DeleteUser(id int)
}

type service_User struct {
	repo repositories.Repo_Users
}

func New_ServiceUser(repo repositories.Repo_Users) Service_User {
	return &service_User{repo: repo}
}

func (s *service_User) RegisterUser(data *models.UserModel) (*models.UserModel, error) {
	created, err := s.repo.RegisterUser(data)
	return created, err
}
func (s *service_User) GetUsers() ([]models.UserModel, error) {
	result, err := s.repo.GetUsers()
	if err != nil {
		return nil, err
	}

	if len(result) == 0 {
		return nil, fmt.Errorf("tidak ada data")
	}

	return result, nil
}
func (s *service_User) GetUserById(id int) (*models.UserModel, error) {
	return s.repo.GetUserById(id)
}
func (s *service_User) EditUser(data *models.UserModel) error {
	return s.repo.EditUser(data)
}
func (s *service_User) DeleteUser(id int) {

}
