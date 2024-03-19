package repositories

import (
	"context"
	"time"

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

func (r *repo_Task) EditTasks(data *models.TasksModel) error {
	tx, err := r.BeginTxx(context.Background(), nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	query := `update tasks set user_id=$1,title=$2,description=$3,status=$4,updated_at=$5 where id=$6`
	_, err = tx.Exec(query, data.UserID, data.Title, data.Description, data.Status, time.Now(), data.ID)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (r *repo_Task) DeleteTasks(id int) error {
	tx, err := r.BeginTxx(context.Background(), nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	query := `delete from tasks where id=$1`
	_, err = tx.Exec(query, id)
	if err != nil {
		return err
	}

	return tx.Commit()
}
