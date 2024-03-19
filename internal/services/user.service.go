package services

import (
	"github.com/funukonta/task_manager/internal/models"
	"github.com/funukonta/task_manager/internal/repositories"
)

type Service_User interface {
	RegisterUser(*models.UserModel) (*models.UserModel, error)
	GetUsers()
	GetUserById(id int)
	EditUser(id int)
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
func (s *service_User) GetUsers() {

}
func (s *service_User) GetUserById(id int) {

}
func (s *service_User) EditUser(id int) {

}
func (s *service_User) DeleteUser(id int) {

}
