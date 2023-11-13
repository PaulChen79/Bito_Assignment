package service

import (
	"bito-assignment/config"
	"bito-assignment/domain"
	"bito-assignment/internal/repo"
	"testing"
)

func TestAddSinglePersonAndMatch(t *testing.T) {
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

	newUser := domain.User{Name: "Sandra", Height: 177, Gender: domain.Female, RemainingDates: 5}

	matches, errFormat := svc.AddSinglePersonAndMatch(&newUser)
	if errFormat != nil {
		t.Errorf("Expected err format nil, but got: %v", errFormat)
	}

	if len(matches) != 1 {
		t.Errorf("Expected 1 user, got %d", len(matches))
	}

	if matches[0].Name != "Bob" {
		t.Errorf("Expected match user name Bob, got %v", matches[0].Name)
	}

	newUser2 := domain.User{Name: "Jason", Height: 168, Gender: domain.Male, RemainingDates: 3}

	matches2, errFormat := svc.AddSinglePersonAndMatch(&newUser2)
	if errFormat != nil {
		t.Errorf("Expected err format nil, but got: %v", errFormat)
	}

	if len(matches2) != 1 {
		t.Errorf("Expected 1 user, got %d", len(matches2))
	}

	if matches2[0].Name != "Alice" {
		t.Errorf("Expected match user name Alice, got %v", matches2[0].Name)
	}

	userSandra := svc.repo.GetUserByName("Sandra")
	if userSandra == nil {
		t.Error("Expected user Sandra, got noting")
	}

	if userSandra.RemainingDates != 4 {
		t.Errorf("Expected user Sandra remaining dates 4, got %d", userSandra.RemainingDates)
	}

	userBob := svc.repo.GetUserByName("Bob")
	if userBob != nil {
		t.Errorf("Expected user Bob to be deleted, got %s", userBob.Name)
	}
}
