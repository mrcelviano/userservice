package logic

import "github.com/mrcelviano/userservice/internal/app"

type userLogic struct {
	repository app.UserRepository
}

func NewUserLogic(repo app.UserRepository) app.UserLogic {
	return &userLogic{repository: repo}
}
