package repo

import (
	"bito-assignment/domain"
	"sync"
)

var (
	userPool = map[string]*domain.User{}
)

type repository struct {
	UserPool *map[string]*domain.User
	Mutex    sync.Mutex
}

func NewRepository() domain.Repository {
	return &repository{
		UserPool: &userPool,
	}
}

func (repo *repository) InitUserPool() {
	*repo.UserPool = map[string]*domain.User{}
}

func (repo *repository) GetUserPool() map[string]*domain.User {
	return *repo.UserPool
}

func (repo *repository) AddUser(user domain.User) {
	repo.Mutex.Lock()
	(*repo.UserPool)[user.Name] = &user
	repo.Mutex.Unlock()
}

func (repo *repository) DeleteUser(user domain.User) {
	repo.Mutex.Lock()
	delete(*repo.UserPool, user.Name)
	repo.Mutex.Unlock()
}

func (repo *repository) GetUserByName(userName string) *domain.User {
	if user, ok := (*repo.UserPool)[userName]; ok {
		return user
	}
	return nil
}
