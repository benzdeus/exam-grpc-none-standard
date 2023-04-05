package controllers

import "gate-way/pkg/pb"

type controller struct {
	AuthService pb.AuthClient
}

func New(a pb.AuthClient) *controller {
	return &controller{
		AuthService: a,
	}
}
