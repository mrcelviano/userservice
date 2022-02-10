package repository

import (
	"context"
	"github.com/mrcelviano/userservice/commons"
	"github.com/mrcelviano/userservice/internal/app"
)

type userRepository struct {
	tagMap     map[string]string
	retCols    []string
	createCols []string
}

func NewUserRepository() app.UserRepository {
	return &userRepository{
		tagMap:     commons.MakeJsonToDbTagMap(app.User{}),
		retCols:    commons.FindDbTags(app.User{}),
		createCols: commons.FindDbTags(app.User{}, "id"),
	}
}

func (u *userRepository) Create(ctx context.Context, user app.User) (resp app.User, err error) {
	sess := commons.DBSessionFromContext(ctx)
	err = sess.
		InsertInto("users").
		Columns(u.createCols...).
		Record(user).
		Returning(u.retCols...).
		LoadContext(ctx, &resp)
	return
}

func (u *userRepository) Update(ctx context.Context, user app.User) (resp app.User, err error) {
	sess := commons.DBSessionFromContext(ctx)

	err = sess.
		Update("users").
		Set("email", user.Email).
		Set("name", user.Name).
		Where("id = ?", user.ID).
		Returning(u.retCols...).
		LoadContext(ctx, &resp)
	return
}

func (u *userRepository) GetByID(ctx context.Context, id int64) (resp app.User, err error) {
	sess := commons.DBSessionFromContext(ctx)

	_, err = sess.
		Select(u.retCols...).
		From("users").
		Where("id = ? and is_deleted = false", id).
		LoadContext(ctx, &resp)
	return
}

func (u *userRepository) GetList(ctx context.Context, p app.Pagination) (resp app.Users, err error) {
	sess := commons.DBSessionFromContext(ctx)

	selectSmt := sess.
		Select(u.retCols...).
		From("users").
		Offset(p.Offset).
		OrderDir(p.SortKey, p.Asc()).
		Where("is_deleted is false")
	p.WithLimit(selectSmt)
	_, err = selectSmt.LoadContext(ctx, &resp)
	return
}

func (u *userRepository) Delete(ctx context.Context, id int64) (err error) {
	sess := commons.DBSessionFromContext(ctx)

	_, err = sess.
		Update("users").
		Set("is_deleted", true).
		Where("id = ?", id).
		ExecContext(ctx)
	return
}

func (u *userRepository) GetTotal(ctx context.Context) (total int64, err error) {
	sess := commons.DBSessionFromContext(ctx)

	_, err = sess.Select("count(*)").
		From("users").
		Where("is_deleted is false").
		LoadContext(ctx, &total)
	return
}

func (u *userRepository) Check(ctx context.Context, user app.User) (isExist bool, err error) {
	sess := commons.DBSessionFromContext(ctx)

	sql := "select exists (select email, name from users where (email = ? or name = ?) and is_deleted = false)"
	_, err = sess.SelectBySql(sql, user.Email, user.Name).LoadContext(ctx, &isExist)
	return
}
