package service

import (
	"bito-assignment/domain"
)

func (svc *service) AddSinglePersonAndMatch(newUser *domain.User) ([]*domain.User, *domain.ErrorFormat) {
	var matches []*domain.User

	userPool := svc.repo.GetUserPool()

	if _, ok := userPool[newUser.Name]; ok {
		return nil, &domain.ErrUserAlreadyExists
	}

	svc.repo.AddUser(*newUser)

	for _, u := range userPool {
		if u.Gender != newUser.Gender &&
			((u.Gender == domain.Male && u.Height > newUser.Height) ||
				(u.Gender == domain.Female && u.Height < newUser.Height)) {

			u.RemainingDates--

			newU := svc.repo.GetUserByName(newUser.Name)
			newU.RemainingDates--

			matches = append(matches, u)

			if u.RemainingDates <= 0 {
				svc.repo.DeleteUser(*u)
			}

			if newUser.RemainingDates <= 0 {
				svc.repo.DeleteUser(*newUser)
			}
		}
	}

	return matches, nil
}
