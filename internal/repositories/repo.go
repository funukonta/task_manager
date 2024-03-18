package repositories

type Repo_UsersInterface interface {
	RegisterUser()
	GetUsers()
	GetUsersById(id int)
	EditUser(id int)
	DeleteUser(id int)
}
