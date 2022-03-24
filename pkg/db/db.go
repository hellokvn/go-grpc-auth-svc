package db

import (
	"log"

	"github.com/hellokvn/go-grpc-auth-svc/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func Init() Handler {
	dbURL := "postgres://kevin@localhost:5432/tasks"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Auth{})

	return Handler{db}
}
