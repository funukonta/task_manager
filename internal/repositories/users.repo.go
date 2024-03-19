package repositories

import (
	"context"
	"time"

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

func (r *repo_Users) GetUserById(id int) (*models.UserModel, error) {
	query := `select * from users where id=$1`
	user := &models.UserModel{}
	err := r.Get(user, query, id)
	return user, err
}

func (r *repo_Users) EditUser(data *models.UserModel) error {
	tx, err := r.BeginTxx(context.Background(), nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	query := `update users set name=$1,email=$2,password=$3,updated_at=$4 where id=$5`
	_, err = tx.Exec(query, data.Name, data.Email, data.Password, time.Now(), data.ID)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (r *repo_Users) DeleteUser(id int) error {
	tx, err := r.BeginTxx(context.Background(), nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	query := `delete from users where id=$1`
	_, err = tx.Exec(query, id)
	if err != nil {
		return err
	}

	return tx.Commit()
}
