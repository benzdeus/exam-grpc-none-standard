package controllers

import (
	"auth-service/pkg/pb"
	"auth-service/pkg/storer"
	"auth-service/pkg/usecases"
	"context"
	"log"
)

// grpc controller
func (c controller) CheckPermission(ctx context.Context, req *pb.CheckPermissionRequest) (*pb.CheckPermissionResponse, error) {
	_storer := storer.NewStore(c.db)
	_usecases := usecases.New(_storer)

	claim, err := _usecases.DecodeToken(req.Token)

	if err != nil {
		return &pb.CheckPermissionResponse{
			Code:    401,
			Message: err.Error(),
		}, nil
	}

	log.Printf("Claim: %v", claim)

	permission, err := _usecases.CheckPermission(claim.Username)
	if err != nil {
		return &pb.CheckPermissionResponse{
			Code:    500,
			Message: err.Error(),
		}, nil
	}

	return &pb.CheckPermissionResponse{
		Permission: permission,
		Code:       200,
		Message:    "success",
	}, nil
}
