package domain

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
