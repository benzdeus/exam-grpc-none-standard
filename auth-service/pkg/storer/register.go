package storer

import "auth-service/pkg/entities"

func (s store) Register(user *entities.User) error {
	return s.db.Create(&user).Error
}
