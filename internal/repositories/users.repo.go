package repositories

import (
	"github.com/jmoiron/sqlx"
)

type Repo_User struct {
	*sqlx.DB
}

func New_RepoUser(db *sqlx.DB) Repo_UsersInterface {
	return &Repo_User{db}
}

func (r *Repo_User) RegisterUser() {

}

func (r *Repo_User) GetUsers() {

}

func (r *Repo_User) GetUsersById(id int) {

}

func (r *Repo_User) EditUser(id int) {

}

func (r *Repo_User) DeleteUser(id int) {

}
