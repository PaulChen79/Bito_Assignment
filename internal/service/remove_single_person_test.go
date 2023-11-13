package service

import (
	"bito-assignment/config"
	"bito-assignment/domain"
	"bito-assignment/internal/repo"
	"testing"
)

func TestRemoveSinglePerson(t *testing.T) {
	svc := NewServiceForTest(&config.Config{}, repo.NewRepository())

	defaultUsers := []domain.User{
		{
			Name:           "Bob",
			Height:         180,
			Gender:         domain.Male,
			RemainingDates: 1,
		},
		{
			Name:           "Alice",
			Height:         160,
			Gender:         domain.Female,
			RemainingDates: 3,
		},
		{
			Name:           "Paul",
			Height:         176,
			Gender:         domain.Male,
			RemainingDates: 10,
		},
		{
			Name:           "Joan",
			Height:         170,
			Gender:         domain.Female,
			RemainingDates: 7,
		},
	}

	for _, u := range defaultUsers {
		svc.repo.AddUser(u)
	}

	errFormat := svc.RemoveSinglePerson("Bob")
	if errFormat != nil {
		t.Errorf("Expected err format nil, but got: %v", errFormat)
	}

	testBob := svc.repo.GetUserByName("Bob")
	if testBob != nil {
		t.Errorf("Expected user Bob to be deleted, but got: %v", testBob)
	}

	errFormat = svc.RemoveSinglePerson("Bob")
	if errFormat == nil {
		t.Error("Expected user Bob to be deleted, but got no err")
	}
}
