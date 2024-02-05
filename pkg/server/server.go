package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ErwinSalas/go-grpc-auth-svc/pkg/auth"
	pb "github.com/ErwinSalas/go-grpc-auth-svc/pkg/proto"
)

type AuthServer struct {
	pb.UnimplementedAuthServiceServer
	auth.AuthService
}

func NewAuthServer(service auth.AuthService) AuthServer
func (s *AuthServer) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	fmt.Println("RPC auth-service/Register")
	result, err := s.AuthService.Register(ctx, req)

	if err != nil {
		return &pb.RegisterResponse{
			Status: http.StatusConflict,
			Error:  "E-Mail already exists",
		}, nil
	}

	return result, nil
}

func (s *AuthServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	fmt.Println("RPC auth-service/Login")
	result, err := s.AuthService.Login(ctx, req)
	if err != nil {
		return &pb.LoginResponse{
			Status: http.StatusNotFound,
			Error:  "User not found",
		}, nil
	}

	return result, nil
}

func (s *AuthServer) Validate(ctx context.Context, req *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	result, err := s.AuthService.Validate(ctx, req)
	if err != nil {
		return &pb.ValidateResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}, nil
	}

	return result, nil
}
