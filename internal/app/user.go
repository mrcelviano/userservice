package app

import "context"

type Users []User

type User struct {
	ID    int64  `json:"id" db:"id"`
	Email string `json:"email" db:"email"`
	Name  string `json:"name" db:"name"`
}

type UserLogic interface {
	Create(context.Context, User) (User, error)
	Update(context.Context, User) (User, error)
	GetByID(context.Context, int64) (User, error)
	GetList(context.Context, Pagination) (Users, error)
	Delete(context.Context, int64) error
}

type UserRepository interface {
	Create(context.Context, User) (User, error)
	Update(context.Context, User) (User, error)
	GetByID(context.Context, int64) (User, error)
	GetList(context.Context, Pagination) (Users, error)
	Delete(context.Context, int64) error
	Check(context.Context, string, string) (bool, bool, error)
}
