package service

import "bito-assignment/domain"

func (svc *service) RemoveSinglePerson(userName string) *domain.ErrorFormat {
	user := svc.repo.GetUserByName(userName)
	if user == nil {
		return &domain.ErrUserNotFound
	}

	svc.repo.DeleteUser(*user)

	return nil
}
