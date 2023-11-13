package service

import (
	"bito-assignment/config"
	"bito-assignment/domain"
	"bito-assignment/internal/repo"
	"testing"
)

func TestQuerySinglePerson(t *testing.T) {
	svc := NewServiceForTest(&config.Config{}, repo.NewRepository())

	defaultUsers := []domain.User{
		{
			Name:           "Bob",
			Height:         180,
			Gender:         domain.Male,
			RemainingDates: 1,
		},
		{
			Name:           "Ken",
			Height:         178,
			Gender:         domain.Male,
			RemainingDates: 12,
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
		{
			Name:           "Sandra",
			Height:         177,
			Gender:         domain.Female,
			RemainingDates: 5,
		},
		{
			Name:           "Guan",
			Height:         190,
			Gender:         domain.Female,
			RemainingDates: 44,
		},
	}

	for _, u := range defaultUsers {
		svc.repo.AddUser(u)
	}

	matches, errFormat := svc.QuerySinglePerson("Sandra", 1)
	if errFormat != nil {
		t.Errorf("Expected err format nil, but got: %v", errFormat)
	}

	if len(matches) != 1 {
		t.Errorf("Expected 1 user, got %d", len(matches))
	}

	for _, u := range matches {
		if u.Name != "Bob" && u.Name != "Ken" {
			t.Errorf("Expected match user name Bob or Ken, got %v", u.Name)
		}
	}

	userBob := svc.repo.GetUserByName("Bob")
	if userBob != nil {
		t.Errorf("Expected user Bob to be deleted, but got %v", userBob)
	}

	userSandra := svc.repo.GetUserByName("Sandra")
	if userSandra == nil {
		t.Errorf("Expected user Sandra to be exist, but got %v", userSandra)
	}

	if userSandra.RemainingDates != 3 {
		t.Errorf("Expected user Sandra remaining dates to be 4, but got %v", userSandra.RemainingDates)
	}

	matches2, errFormat := svc.QuerySinglePerson("Ken", 2)
	if errFormat != nil {
		t.Errorf("Expected err format nil for match 2, but got: %v", errFormat)
	}

	if len(matches2) != 2 {
		t.Errorf("Expected 2 user, got %d", len(matches2))
	}

	for _, u := range matches2 {
		if u.Name != "Alice" && u.Name != "Joan" && u.Name != "Sandra" {
			t.Errorf("Expected match user name Alice or Joan or Sandra, got %v", u.Name)
		}
	}

	matches3, errFormat := svc.QuerySinglePerson("Ken", 5)
	if errFormat != nil {
		t.Errorf("Expected err format nil for match 3, but got: %v", errFormat)
	}

	if len(matches3) != 3 {
		t.Errorf("Expected 3 user, got %d", len(matches3))
	}
}
