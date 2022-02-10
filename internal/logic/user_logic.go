package logic

import (
	"context"
	"github.com/mrcelviano/userservice/internal/app"
	"github.com/pkg/errors"
	"log"
)

type userLogic struct {
	repository   app.UserRepository
	notification app.NotificationGRPCRepository
}

func NewUserLogic(repo app.UserRepository, notification app.NotificationGRPCRepository) app.UserLogic {
	return &userLogic{
		repository:   repo,
		notification: notification,
	}
}

func (u *userLogic) Create(ctx context.Context, user app.User) (resp app.User, err error) {
	log.Println("check user fields")
	isExist, err := u.repository.Check(ctx, user)
	if err != nil {
		return resp, errors.Wrap(err, "can`t check user from db")
	}
	if isExist {
		return resp, errors.New("email or name already exists in the database")
	}

	resp, err = u.repository.Create(ctx, user)
	if err != nil {
		return resp, errors.Wrap(err, "can`t create user")
	}

	//отправить уведомление в notification service
	_, err = u.notification.SendNotification(ctx, resp)
	if err != nil {
		log.Println("can`t send notification. Error: ", err.Error())
	}
	return
}

func (u *userLogic) Update(ctx context.Context, user app.User) (resp app.User, err error) {
	return u.repository.Update(ctx, user)
}

func (u *userLogic) GetByID(ctx context.Context, id int64) (resp app.User, err error) {
	return u.repository.GetByID(ctx, id)
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
	return u.repository.Delete(ctx, id)
}
