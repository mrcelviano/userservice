package logic

import (
	"context"
	"github.com/mrcelviano/userservice/internal/app"
)

type userLogic struct {
	repository app.UserRepository
}

func NewUserLogic(repo app.UserRepository) app.UserLogic {
	return &userLogic{repository: repo}
}

func (u *userLogic) Create(ctx context.Context, user app.User) (resp app.User, err error) {
	return
}

func (u *userLogic) Update(ctx context.Context, user app.User) (resp app.User, err error) {
	return
}

func (u *userLogic) GetByID(ctx context.Context, id int64) (resp app.User, err error) {
	return
}

func (u *userLogic) GetList(ctx context.Context, p app.Pagination) (resp app.Users, err error) {
	return
}

func (u *userLogic) Delete(ctx context.Context, id int64) (err error) {
	return
}
