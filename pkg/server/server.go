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

func NewAuthServer(service auth.AuthService) AuthServer
func (s *AuthServer) Register(ctx context.Context, req *authpb.RegisterRequest) (*authpb.RegisterResponse, error) {
	fmt.Println("RPC auth-service/Register")
	result, err := s.AuthService.Register(ctx, req)

	if err != nil {
		return &authpb.RegisterResponse{
			Status: http.StatusConflict,
			Error:  "E-Mail already exists",
		}, nil
	}

	return result, nil
}

func (s *AuthServer) Login(ctx context.Context, req *authpb.LoginRequest) (*authpb.LoginResponse, error) {
	fmt.Println("RPC auth-service/Login")
	result, err := s.AuthService.Login(ctx, req)
	if err != nil {
		return &authpb.LoginResponse{
			Status: http.StatusNotFound,
			Error:  "User not found",
		}, nil
	}

	return result, nil
}

func (s *AuthServer) Validate(ctx context.Context, req *authpb.ValidateRequest) (*authpb.ValidateResponse, error) {
	result, err := s.AuthService.Validate(ctx, req)
	if err != nil {
		return &authpb.ValidateResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}, nil
	}

	return result, nil
}
