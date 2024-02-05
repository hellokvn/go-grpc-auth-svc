package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ErwinSalas/go-grpc-auth-svc/pkg/auth"
	authpb "github.com/ErwinSalas/go-grpc-auth-svc/pkg/proto"
)

type AuthServer struct {
	authpb.UnimplementedAuthServiceServer
	auth.AuthService
}

func NewAuthServer(service auth.AuthService) authpb.AuthServiceServer {
	return &AuthServer{
		authpb.UnimplementedAuthServiceServer{},
		service,
	}
}

func (s *AuthServer) Register(ctx context.Context, req *authpb.RegisterRequest) (*authpb.RegisterResponse, error) {
	fmt.Println("RPC auth-service/Register")
	result, err := s.AuthService.Register(ctx, req)

	if err != nil {
		return &authpb.RegisterResponse{
			Status: http.StatusConflict,
			Error:  "E-Mail already exists",
		}, err
	}

	return result, nil
}

func (s *AuthServer) Login(ctx context.Context, req *authpb.LoginRequest) (*authpb.LoginResponse, error) {
	fmt.Println("RPC auth-service/Login")
	result, err := s.AuthService.Login(ctx, req)
	if err != nil {
		fmt.Println(err.Error())
		return &authpb.LoginResponse{
			Status: http.StatusNotFound,
			Error:  "User not found",
		}, err
	}

	return result, nil
}

func (s *AuthServer) Validate(ctx context.Context, req *authpb.ValidateRequest) (*authpb.ValidateResponse, error) {
	result, err := s.AuthService.Validate(ctx, req)
	if err != nil {
		fmt.Println(err.Error())
		return &authpb.ValidateResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}, err
	}

	return result, nil
}

func (s *AuthServer) HealthCheck(ctx context.Context, req *authpb.Empty) (*authpb.Empty, error) {
	return &authpb.Empty{}, nil
}
