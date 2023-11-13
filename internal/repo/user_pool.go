package repo

import (
	"bito-assignment/domain"
)

var (
	userPool = map[string]User{}
)

type repository struct {
}

func NewRepository() domain.Repository {
	return &repository{}
}

type Gender string

const (
	Male   Gender = "male"
	Female Gender = "female"
)

type User struct {
	Name           string
	Height         int
	Gender         Gender
	RemainingDates int
}
