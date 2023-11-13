package domain

type Service interface {
	AddSinglePersonAndMatch(newUser *User) ([]*User, *ErrorFormat)
}
