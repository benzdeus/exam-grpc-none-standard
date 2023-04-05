package controllers

import (
	"auth-service/pkg/pb"

	"gorm.io/gorm"
)

type controller struct {
	// usecase entities.UserUsecase
	db *gorm.DB
	pb.UnsafeAuthServer
}

func New(db *gorm.DB) *controller {
	return &controller{
		db: db,
	}
}
