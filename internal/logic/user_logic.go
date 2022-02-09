package logic

import (
	"context"
	"github.com/mrcelviano/userservice/internal/app"
	"github.com/pkg/errors"
	"log"
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
	return u.Update(ctx, user)
}

func (u *userLogic) GetByID(ctx context.Context, id int64) (resp app.User, err error) {
	return u.GetByID(ctx, id)
}

func (u *userLogic) GetList(ctx context.Context, p app.Pagination) (resp app.PaginationUsers, err error) {
	log.Println("getting user list")
	resp.Result, err = u.repository.GetList(ctx, p)
	if err != nil {
		return resp, errors.Wrap(err, "cant get list")
	}

	log.Println("get total ef")
	resp.Total, err = u.repository.GetTotal(ctx)
	if err != nil {
		return resp, errors.Wrap(err, "cant get total users")
	}
	return
}

func (u *userLogic) Delete(ctx context.Context, id int64) (err error) {
	return u.Delete(ctx, id)
}
