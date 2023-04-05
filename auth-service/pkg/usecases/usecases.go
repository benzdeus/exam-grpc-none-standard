package usecases

import "auth-service/pkg/entities"

type usecase struct {
	repo entities.UserRepo
}

func New(repo entities.UserRepo) *usecase {
	return &usecase{repo: repo}
}
