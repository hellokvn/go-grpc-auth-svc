package main

import (
	"fmt"
	"log"
	"net"

	"github.com/hellokvn/go-grpc-auth-svc/pkg/config"
	"github.com/hellokvn/go-grpc-auth-svc/pkg/db"
	"github.com/hellokvn/go-grpc-auth-svc/pkg/pb"
	"github.com/hellokvn/go-grpc-auth-svc/pkg/services"
	"github.com/hellokvn/go-grpc-auth-svc/pkg/utils"
	"google.golang.org/grpc"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	h := db.Init(c.DBUrl)

	jwt := utils.JwtWrapper{
		SecretKey:       c.JWTSecretKey,
		Issuer:          "go-grpc-auth-svc",
		ExpirationHours: 24 * 365,
	}

	lis, err := net.Listen("tcp", c.Port)

	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	fmt.Println("Auth Svc on", c.Port)

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
