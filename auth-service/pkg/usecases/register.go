package usecases

import "auth-service/pkg/entities"

func (u *usecase) Register(user *entities.User) error {
	return u.repo.Register(user)
}
