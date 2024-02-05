package main

import (
	"fmt"
	"log"
	"net"

	"github.com/ErwinSalas/go-grpc-auth-svc/pkg/auth"
	"github.com/ErwinSalas/go-grpc-auth-svc/pkg/config"
	"github.com/ErwinSalas/go-grpc-auth-svc/pkg/database"
	authpb "github.com/ErwinSalas/go-grpc-auth-svc/pkg/proto"
	"github.com/ErwinSalas/go-grpc-auth-svc/pkg/server"
	"github.com/ErwinSalas/go-grpc-auth-svc/pkg/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	h := database.Init(c.DBUrl)

	jwt := utils.JwtWrapper{
		SecretKey:       c.JWTSecretKey,
		Issuer:          "go-grpc-auth-svc",
		ExpirationHours: 24 * 365,
	}

	listen, err := net.Listen("tcp", c.Port)

	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	fmt.Println("Auth Svc on", c.Port)

	grpcServer := grpc.NewServer()
	authService := auth.NewAuthService(auth.NewUserRepository(h), jwt) // Puedes pasar una conexión de base de datos real aquí.
	authpb.RegisterAuthServiceServer(grpcServer, server.NewAuthServer(authService))

	// Register reflection service on gRPC server.
	reflection.Register(grpcServer)

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
