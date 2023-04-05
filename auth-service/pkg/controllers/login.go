package controllers

import (
	"auth-service/pkg/pb"
	"auth-service/pkg/storer"
	"auth-service/pkg/usecases"
	"context"
)

// grpc controller
func (c controller) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	_storer := storer.NewStore(c.db)
	_usecases := usecases.New(_storer)

	user, err := _usecases.Login(req.Username, req.Password)
	if err != nil {
		return nil, err
	}

	token, err := _usecases.EncodeToken(*user)

	if err != nil {
		return &pb.LoginResponse{
			Code:    500,
			Message: err.Error(),
		}, nil
	}

	return &pb.LoginResponse{
		Token: token,
	}, nil
}
