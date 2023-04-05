package storer

import "auth-service/pkg/entities"

func (s store) CheckPermission(username string) (string, error) {
	user := entities.User{}
	err := s.db.Debug().Select([]string{"permission"}).Where("username=?", username).First(&user).Error

	if err != nil {
		return "", err
	}

	return user.Permission, nil
}
