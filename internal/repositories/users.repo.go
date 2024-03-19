package repositories

import (
	"context"

	"github.com/funukonta/task_manager/internal/models"
	"github.com/jmoiron/sqlx"
)

type repo_Users struct {
	*sqlx.DB
}

func New_RepoUser(db *sqlx.DB) Repo_Users {
	return &repo_Users{db}
}

func (r *repo_Users) RegisterUser(data *models.UserModel) (*models.UserModel, error) {
	tx, err := r.BeginTxx(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	query := `insert into users (name,email,password) values ($1,$2,$3)
	returning *`
	userCreated := models.UserModel{}
	if err := tx.Get(&userCreated, query, data.Name, data.Email, data.Password); err != nil {
		return nil, err
	}

	err = tx.Commit()
	return &userCreated, err
}

func (r *repo_Users) GetUsers() ([]models.UserModel, error) {
	result := []models.UserModel{}
	query := `select * from users`
	err := r.Select(&result, query)
	return result, err
}

func (r *repo_Users) GetUserById(id int) {

}

func (r *repo_Users) EditUser(id int) {

}

func (r *repo_Users) DeleteUser(id int) {

}
