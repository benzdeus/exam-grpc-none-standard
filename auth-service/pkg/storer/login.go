package storer

import "auth-service/pkg/entities"

func (s store) Login(username, password string) (*entities.User, error) {
	user := entities.User{}
	err := s.db.
		Select([]string{"id", "username", "permission"}).
		First(&user, "username = ? AND password=?", username, password).
		Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}
