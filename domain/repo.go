package domain

type Repository interface {
	InitUserPool()
	GetUserPool() map[string]*User
	AddUser(user User)
	DeleteUser(user User)
	GetUserByName(userName string) *User
}
