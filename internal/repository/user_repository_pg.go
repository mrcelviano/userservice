package repository

import (
	"context"
	goCraft "github.com/gocraft/dbr"
	"github.com/mrcelviano/userservice/internal/domain"
	"github.com/mrcelviano/userservice/pkg/database/postgres"
)

type userRepositoryPG struct {
	pgSession *goCraft.Session
}

func NewUserRepositoryPG(pgDBConn *goCraft.Connection) domain.UserRepositoryPG {
	return &userRepositoryPG{
		pgSession: pgDBConn.NewSession(&eventReceiver{}),
	}
}

func (u *userRepositoryPG) Create(ctx context.Context, user domain.User) (domain.User, error) {
	var resp domain.User
	err := u.pgSession.
		InsertInto("users").
		Columns("email", "name").
		Record(user).
		Returning("id", "email", "name").
		LoadContext(ctx, &resp)
	if err != nil && postgres.GetError(err) == postgres.ErrDuplicateValue {
		return resp, domain.ErrBadParamInput
	}
	return resp, err
}

func (u *userRepositoryPG) Update(ctx context.Context, user domain.User) (domain.User, error) {
	var resp domain.User
	err := u.pgSession.
		Update("users").
		Set("email", user.Email).
		Set("name", user.Name).
		Where("id = ?", user.ID).
		Returning("id", "email", "name").
		LoadContext(ctx, &resp)
	if err != nil && postgres.GetError(err) == postgres.ErrDuplicateValue {
		return resp, domain.ErrBadParamInput
	}
	return resp, err
}

func (u *userRepositoryPG) GetByID(ctx context.Context, id int64) (domain.User, error) {
	var resp domain.User
	_, err := u.pgSession.
		Select("id", "email", "name").
		From("users").
		Where("id = ? and is_deleted = false", id).
		LoadContext(ctx, &resp)
	if err != nil && err == goCraft.ErrNotFound {
		return resp, domain.ErrNotFound
	}
	return resp, err
}

func (u *userRepositoryPG) GetList(ctx context.Context, p domain.GetUserListRequest) ([]domain.User, error) {
	var resp []domain.User
	selectSmt := u.pgSession.
		Select("id", "email", "name").
		From("users").
		Where("is_deleted is false")

	postgres.SetSortKeyAndSortOrder(p.SortKey, p.SortOrder, selectSmt)
	postgres.SetLimit(p.Limit, selectSmt)
	postgres.SetOffset(p.Offset, selectSmt)

	if _, err := selectSmt.LoadContext(ctx, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (u *userRepositoryPG) Delete(ctx context.Context, id int64) error {
	_, err := u.pgSession.
		Update("users").
		Set("is_deleted", true).
		Where("id = ?", id).
		ExecContext(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (u *userRepositoryPG) GetTotal(ctx context.Context) (int64, error) {
	var total int64
	_, err := u.pgSession.
		Select("count(*)").
		From("users").
		Where("is_deleted is false").
		LoadContext(ctx, &total)
	if err != nil {
		return 0, err
	}
	return total, nil
}
