package repositories

import (
	"context"

	"github.com/funukonta/task_manager/internal/models"
	"github.com/jmoiron/sqlx"
)

type repo_Task struct {
	*sqlx.DB
}

func New_RepoTask(db *sqlx.DB) Repo_Tasks {
	return &repo_Task{db}
}

func (r *repo_Task) RegisterTasks(data *models.TasksModel) (*models.TasksModel, error) {
	tx, err := r.BeginTxx(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	query := `insert tasks (user_id,title,description) values ($1,$2,$3)
	returning *
	`
	result := &models.TasksModel{}
	if err := tx.Get(result, query, data.UserID, data.Title, data.Description); err != nil {
		return nil, err
	}

	err = tx.Commit()

	return result, err
}

func (r *repo_Task) GetTasks() ([]models.TasksModel, error) {
	result := []models.TasksModel{}
	query := `select * from tasks`
	if err := r.Select(&result, query); err != nil {
		return nil, err
	}

	return result, nil
}

func (r *repo_Task) GetTasksById(id int) (*models.TasksModel, error) {
	result := &models.TasksModel{}
	query := `select * from tasks where id=$1`
	if err := r.Get(result, query, id); err != nil {
		return nil, err
	}

	return result, nil
}

func (r *repo_Task) EditTask(id int) {

}

func (r *repo_Task) DeleteTask(id int) {

}
