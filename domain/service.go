package domain

type Service interface {
	AddSinglePersonAndMatch(newUser *User) ([]*User, *ErrorFormat)
	RemoveSinglePerson(userName string) *ErrorFormat
	QuerySinglePerson(userName string, number uint) ([]*User, *ErrorFormat)
}
