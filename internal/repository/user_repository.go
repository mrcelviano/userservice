package repository

import "social-tech/userservice/internal/app"

type userRepository struct{}

func NewUserRepository() app.UserRepository {
	return &userRepository{}
}
