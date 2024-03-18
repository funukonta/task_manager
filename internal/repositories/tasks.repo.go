package repositories

import (
	"github.com/jmoiron/sqlx"
)

type Repo_Task struct {
	*sqlx.DB
}

func New_RepoTask(db *sqlx.DB) Repo_TaksInterface {
	return &Repo_Task{db}
}

func (r *Repo_Task) AddTask() {

}

func (r *Repo_Task) GetTasks() {

}

func (r *Repo_Task) GetTaskById(id int) {

}

func (r *Repo_Task) EditTask(id int) {

}

func (r *Repo_Task) DeleteTask(id int) {

}
