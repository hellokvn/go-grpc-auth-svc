package main

import (
	"fmt"
	"log"
	"net"

	"github.com/hellokvn/go-grpc-auth-svc/pkg/db"
	"github.com/hellokvn/go-grpc-auth-svc/pkg/pb"
	"github.com/hellokvn/go-grpc-auth-svc/pkg/services"
	"github.com/hellokvn/go-grpc-auth-svc/pkg/utils"
	"google.golang.org/grpc"
)

func main() {
	port := ":50051"
	h := db.Init("postgres://kevin@localhost:5432/auth_svc")

	jwt := utils.JwtWrapper{
		SecretKey:       "SECRET_KEY",
		Issuer:          "go-grpc-auth-svc",
		ExpirationHours: 24 * 365,
	}

	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	fmt.Println("Auth Svc on", port)

	s := services.Server{
		H:   h,
		Jwt: jwt,
	}

	grpcServer := grpc.NewServer()

	pb.RegisterAuthServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
