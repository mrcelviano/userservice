package repository

import "github.com/mrcelviano/userservice/internal/app"

type userRepository struct{}

func NewUserRepository() app.UserRepository {
	return &userRepository{}
}
