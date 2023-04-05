package utils

import (
	"context"
	"gate-way/pkg/pb"
	"log"
)

var AuthService pb.AuthClient

func CheckPermission(token string) string {
	res, err := AuthService.CheckPermission(context.Background(), &pb.CheckPermissionRequest{
		Token: token,
	})

	if err != nil {
		return ""
	}

	log.Printf("%#+v", res)

	return res.Permission
}
