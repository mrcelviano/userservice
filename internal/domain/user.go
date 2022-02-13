package domain

import (
	"context"
	"github.com/mrcelviano/userservice/pkg/validator"
)

type GetUserListResponse struct {
	Total  int64  `json:"total"`
	Result []User `json:"result"`
}

type GetUserListRequest struct {
	Limit     uint64 `query:"limit"`
	Offset    uint64 `query:"offset"`
	SortKey   string `query:"sortKey"`
	SortOrder string `query:"sortOrder"`
}

type User struct {
	ID    int64  `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

func (u *User) ValidateUserFields() error {
	isValid, err := validator.ValidateEmail(u.Email)
	if err != nil || !isValid {
		return ErrInvalidEmail
	}

	return nil
}

type UserService interface {
	Create(context.Context, User) (User, error)
	Update(context.Context, User) (User, error)
	GetByID(context.Context, int64) (User, error)
	GetList(context.Context, GetUserListRequest) (GetUserListResponse, error)
	Delete(context.Context, int64) error
}

type UserRepositoryPG interface {
	Create(context.Context, User) (User, error)
	Update(context.Context, User) (User, error)
	GetByID(context.Context, int64) (User, error)
	GetList(context.Context, GetUserListRequest) ([]User, error)
	GetTotal(context.Context) (int64, error)
	Delete(context.Context, int64) error
}
