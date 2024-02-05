package auth

import (
	"fmt"

	"github.com/ErwinSalas/go-grpc-auth-svc/pkg/database"
	"github.com/ErwinSalas/go-grpc-auth-svc/pkg/models"
)

type UserRepository interface {
	CreateUser(user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
}

type userRepository struct {
	datastore database.Handler
}

func NewUserRepository(db database.Handler) UserRepository {
	return &userRepository{datastore: db}
}

func (r *userRepository) CreateUser(user *models.User) error {
	var newUser models.User
	if result := r.datastore.DB.Where(user.Email).First(&newUser); result.Error == nil {
		return fmt.Errorf("E-Mail already exists")
	}

	if err := r.datastore.DB.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

func (r *userRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User

	if result := r.datastore.DB.Where(&models.User{Email: email}).First(&user); result.Error != nil {
		return nil, fmt.Errorf("User not found")
	}

	return &user, nil
}
