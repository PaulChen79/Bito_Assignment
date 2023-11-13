package service

import (
	"bito-assignment/domain"
	"math/rand"
	"time"
)

func (svc *service) QuerySinglePerson(userName string, number uint) ([]*domain.User, *domain.ErrorFormat) {
	user := svc.repo.GetUserByName(userName)
	if user == nil {
		return nil, &domain.ErrUserNotFound
	}

	var matches []*domain.User

	userPool := svc.repo.GetUserPool()

	for _, u := range userPool {
		if u.Gender != user.Gender &&
			((u.Gender == domain.Male && u.Height > user.Height) ||
				(u.Gender == domain.Female && u.Height < user.Height)) {

			u.RemainingDates--
			user.RemainingDates--

			matches = append(matches, u)

			if u.RemainingDates <= 0 {
				svc.repo.DeleteUser(*u)
			}

			if user.RemainingDates <= 0 {
				svc.repo.DeleteUser(*user)
			}
		}
	}

	rand.New(rand.NewSource(time.Now().UnixNano()))

	rand.Shuffle(len(matches), func(i, j int) {
		matches[i], matches[j] = matches[j], matches[i]
	})

	if len(matches) <= int(number) {
		return matches, nil
	}

	// 否則，只返回前 N 個元素
	return matches[:number], nil
}
