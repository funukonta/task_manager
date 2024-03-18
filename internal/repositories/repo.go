package repositories

type Repo_UsersInterface interface {
	RegisterUser()
	GetUsers()
	GetUserById(id int)
	EditUser(id int)
	DeleteUser(id int)
}

type Repo_TaksInterface interface {
	AddTask()
	GetTasks()
	GetTaskById(id int)
	EditTask(id int)
	DeleteTask(id int)
}
