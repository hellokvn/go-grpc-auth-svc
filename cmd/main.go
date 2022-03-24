package main

import (
	"fmt"
	"log"
	"net"

	"github.com/hellokvn/go-grpc-auth-svc/pkg/db"
	"github.com/hellokvn/go-grpc-auth-svc/pkg/pb"
	"github.com/hellokvn/go-grpc-auth-svc/pkg/service"
	"github.com/hellokvn/go-grpc-auth-svc/pkg/utils"
	"google.golang.org/grpc"
)

func main() {
	port := ":50052"
	h := db.Init()

	jwt := utils.JwtWrapper{
		SecretKey:       "SECRET_KEY",
		Issuer:          "go-grpc-auth-svc",
		ExpirationHours: 24 * 365,
	}

	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	fmt.Println("Run Product Svc on", port)

	s := service.Server{
		H:   h,
		Jwt: jwt,
	}

	grpcServer := grpc.NewServer()

	pb.RegisterAuthServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
