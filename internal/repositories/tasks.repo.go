package repositories

import (
	"github.com/jmoiron/sqlx"
)

type repo_Task struct {
	*sqlx.DB
}

func New_RepoTask(db *sqlx.DB) Repo_Tasks {
	return &repo_Task{db}
}

func (r *repo_Task) AddTask() {

}

func (r *repo_Task) GetTasks() {

}

func (r *repo_Task) GetTaskById(id int) {

}

func (r *repo_Task) EditTask(id int) {

}

func (r *repo_Task) DeleteTask(id int) {

}
