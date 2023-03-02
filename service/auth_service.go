// Package service  @Author xiaobaiio 2023/2/23 10:26:00
package service

import (
	"context"
	"github.com/xiaopangio/pcbook/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthServer struct {
	userStore  UserStore
	jwtManager *JWTManager
	pb.UnimplementedAuthServiceServer
}

func NewAuthServer(userStore UserStore, jwtManager *JWTManager) *AuthServer {
	return &AuthServer{userStore: userStore, jwtManager: jwtManager, UnimplementedAuthServiceServer: pb.UnimplementedAuthServiceServer{}}
}
func (server *AuthServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	username := req.GetUsername()
	password := req.GetPassword()
	user, err := server.userStore.Find(username)
	if err != nil {
		return nil, logError(status.Errorf(codes.Internal, "cannot find user: %v", err))
	}
	if user == nil || !user.IsCorrectPassword(password) {
		return nil, status.Errorf(codes.NotFound, "incorrect username or password")
	}
	token, err := server.jwtManager.Generate(user)
	if err != nil {
		return nil, logError(status.Errorf(codes.Internal, "cannot generate access token: %v", err))
	}
	res := &pb.LoginResponse{
		AccessToken: token,
	}
	return res, nil
}
