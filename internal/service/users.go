package service

import (
	"context"
	"github.com/mrcelviano/userservice/internal/domain"
	"github.com/mrcelviano/userservice/pkg/logger"
	"github.com/mrcelviano/userservice/pkg/notification"
)

type userService struct {
	userPG       domain.UserRepositoryPG
	notification notification.Service
}

func NewUserService(userPG domain.UserRepositoryPG, notification notification.Service) domain.UserService {
	return &userService{
		userPG:       userPG,
		notification: notification,
	}
}

func (u *userService) Create(ctx context.Context, user domain.User) (domain.User, error) {
	logger.Info("check valid email")
	err := user.ValidateUserFields()
	if err != nil {
		return user, domain.ErrInvalidEmail
	}

	newUser, err := u.userPG.Create(ctx, user)
	if err != nil {
		return user, err
	}

	isRegistered, err := u.notification.RegisterNotification(ctx, newUser.ID)
	if err != nil && !isRegistered {
		logger.Errorf("can`t registered notification: %s\n", err.Error())
	}
	return newUser, nil
}

func (u *userService) Update(ctx context.Context, user domain.User) (domain.User, error) {
	resp, err := u.userPG.Update(ctx, user)
	if err != nil {
		logger.Errorf("can`t update user: %s\n", err.Error())
		return resp, domain.ErrInternalServerError
	}
	return resp, nil
}

func (u *userService) GetByID(ctx context.Context, id int64) (domain.User, error) {
	resp, err := u.userPG.GetByID(ctx, id)
	if err != nil {
		logger.Errorf("can`t get user by id: %s\n", err.Error())
		return resp, domain.ErrInternalServerError
	}
	return resp, nil
}

func (u *userService) GetList(ctx context.Context, p domain.GetUserListRequest) (domain.GetUserListResponse, error) {
	logger.Info("getting user list")
	userList, err := u.userPG.GetList(ctx, p)
	if err != nil {
		logger.Errorf("can`t get user list: %s\n", err.Error())
		return domain.GetUserListResponse{}, domain.ErrInternalServerError
	}

	logger.Info("get total ef")
	userTotal, err := u.userPG.GetTotal(ctx)
	if err != nil {
		logger.Errorf("can`t get total user: %s\n", err.Error())
		return domain.GetUserListResponse{}, domain.ErrInternalServerError
	}
	return domain.GetUserListResponse{
		Total:  userTotal,
		Result: userList,
	}, nil
}

func (u *userService) Delete(ctx context.Context, id int64) error {
	err := u.userPG.Delete(ctx, id)
	if err != nil {
		logger.Errorf("can`t delete user: %s\n", err.Error())
		return domain.ErrInternalServerError
	}
	return nil
}

func (u *userService) SetIsRegisteredStatus(ctx context.Context, id int64) error {
	err := u.userPG.SetIsRegisteredStatus(ctx, id)
	if err != nil {
		logger.Errorf("can`t set is registered status: %s\n", err.Error())
		return domain.ErrInternalServerError
	}
	return nil
}
