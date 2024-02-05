package auth

import (
	"context"
	"net/http"

	"github.com/ErwinSalas/go-grpc-auth-svc/pkg/models"
	authpb "github.com/ErwinSalas/go-grpc-auth-svc/pkg/proto"
	"github.com/ErwinSalas/go-grpc-auth-svc/pkg/utils"
)

type AuthService interface {
	Register(ctx context.Context, req *authpb.RegisterRequest) (*authpb.RegisterResponse, error)
	Login(ctx context.Context, req *authpb.LoginRequest) (*authpb.LoginResponse, error)
	Validate(ctx context.Context, req *authpb.ValidateRequest) (*authpb.ValidateResponse, error)
}

type authService struct {
	repo UserRepository
	jwt  utils.JwtWrapper
}

func NewAuthService(repo UserRepository, jwt utils.JwtWrapper) AuthService {
	return &authService{
		repo: repo,
		jwt:  jwt,
	}
}

func (s *authService) Register(ctx context.Context, req *authpb.RegisterRequest) (*authpb.RegisterResponse, error) {
	user, err := s.repo.GetUserByEmail(req.Email)
	if err == nil {
		return &authpb.RegisterResponse{
			Status: http.StatusConflict,
			Error:  "E-Mail already exists",
		}, nil
	}

	user = &models.User{
		Email:    req.Email,
		Password: utils.HashPassword(req.Password),
	}

	if err := s.repo.CreateUser(user); err != nil {
		return nil, err
	}

	return &authpb.RegisterResponse{
		Status: http.StatusCreated,
	}, nil
}

func (s *authService) Login(ctx context.Context, req *authpb.LoginRequest) (*authpb.LoginResponse, error) {
	user, err := s.repo.GetUserByEmail(req.Email)
	if err != nil {
		return &authpb.LoginResponse{
			Status: http.StatusNotFound,
			Error:  "User not found",
		}, nil
	}

	match := utils.CheckPasswordHash(req.Password, user.Password)
	if !match {
		return &authpb.LoginResponse{
			Status: http.StatusUnauthorized,
			Error:  "Invalid password",
		}, nil
	}

	token, err := s.jwt.GenerateToken(*user)
	if err != nil {
		return nil, err
	}

	return &authpb.LoginResponse{
		Status: http.StatusOK,
		Token:  token,
	}, nil
}

func (s *authService) Validate(ctx context.Context, req *authpb.ValidateRequest) (*authpb.ValidateResponse, error) {
	claims, err := s.jwt.ValidateToken(req.Token)
	if err != nil {
		return &authpb.ValidateResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}, nil
	}

	user, err := s.repo.GetUserByEmail(claims.Email)
	if err != nil {
		return &authpb.ValidateResponse{
			Status: http.StatusNotFound,
			Error:  "User not found",
		}, nil
	}

	return &authpb.ValidateResponse{
		Status: http.StatusOK,
		UserId: user.Id,
	}, nil
}
