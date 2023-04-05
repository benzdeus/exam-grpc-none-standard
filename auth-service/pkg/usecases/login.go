package usecases

import "auth-service/pkg/entities"

func (u usecase) Login(username, password string) (*entities.User, error) {
	return u.repo.Login(username, password)
}
