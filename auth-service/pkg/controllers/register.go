package controllers

import (
	"auth-service/pkg/entities"
	"auth-service/pkg/pb"
	"auth-service/pkg/storer"
	"auth-service/pkg/usecases"
	"context"
)

// grpc controller
func (c controller) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {

	user := &entities.User{
		Username:   req.Username,
		Password:   req.Password,
		Permission: req.Permission,
	}

	_storer := storer.NewStore(c.db)
	_usecases := usecases.New(_storer)

	err := _usecases.Register(user)
	if err != nil {
		return &pb.RegisterResponse{
			Code:    500,
			Message: err.Error(),
		}, nil
	}

	return &pb.RegisterResponse{
		Code:    200,
		Message: "success",
	}, nil
}
