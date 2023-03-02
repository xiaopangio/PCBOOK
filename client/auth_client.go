// Package client  @Author xiaobaiio 2023/2/23 15:19:00
package client

import (
	"context"
	"github.com/xiaopangio/pcbook/pb"
	"google.golang.org/grpc"
	"time"
)

type AuthClient struct {
	service  pb.AuthServiceClient
	username string
	password string
}

func NewAuthClient(cc *grpc.ClientConn, username string, password string) *AuthClient {
	return &AuthClient{service: pb.NewAuthServiceClient(cc), username: username, password: password}
}
func (client *AuthClient) Login() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	req := &pb.LoginRequest{
		Username: client.username,
		Password: client.password,
	}
	res, err := client.service.Login(ctx, req)
	if err != nil {
		return "", err
	}
	return res.GetAccessToken(), nil
}
