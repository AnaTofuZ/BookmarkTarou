package mysql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"unsafe"

	"github.com/anatofuz/BookmarkTarou/infra/record"
	"github.com/anatofuz/BookmarkTarou/infra/store"
	"github.com/anatofuz/BookmarkTarou/model"
	"github.com/volatiletech/sqlboiler/boil"
)

type userStoreImpl struct {
	db *sql.DB
}

// NewUserStore UserからDB関連のインスタンス生成
func NewUserStore(db *sql.DB) store.UserStore {
	return &userStoreImpl{db: db}
}

func (u *userStoreImpl) GetUserFromID(ctx context.Context, uid uint64) (*model.User, error) {
	user, err := record.FindUser(ctx, u.db, uid)
	if err != nil {
		if !errors.Is(sql.ErrNoRows, err) {
			return nil, model.UserErrNotFound
		}
		return nil, fmt.Errorf("failed find user: %w at mysql.GetUserFromID", err)
	}

	return &model.User{
		ID:   user.ID,
		Name: string(user.Name),
	}, nil
}

func (u *userStoreImpl) CreateUser(ctx context.Context, name string, password []byte) (*model.User, error) {
	recordUser := &record.User{
		Name:     *(*[]byte)(unsafe.Pointer(&name)),
		Password: password,
	}
	tx, err := u.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed create user: %w at mysql.CreateUser", err)
	}
	err = recordUser.Insert(ctx, tx, boil.Infer())

	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("failed create user: %w at mysql.CreateUser", err)
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("failed create user: %w at mysql.CreateUser", err)
	}
	return &model.User{
		ID:   recordUser.ID,
		Name: string(recordUser.Name),
	}, nil
}
func (u *userStoreImpl) GetPasswordWithUserFromName(ctx context.Context, name string) (*model.UserWithPW, error) {
	user, err := record.Users(qm.Where("name=?", name)).One(ctx, u.db)
	if err != nil {
		return nil, fmt.Errorf("failed GetPasswordFromName: %w", err)
	}
	return &model.UserWithPW{
		User: model.User{ID: user.ID, Name: string(user.Name)},
		Pw:   user.Password,
	}, nil
}

func (u *userStoreImpl) UpdateUser(ctx context.Context, uid uint64, name string, password []byte) (*model.User, error) {
	tx, err := u.db.BeginTx(ctx, nil)
	usr, err := record.FindUser(ctx, tx, uid)
	if err != nil {
		if !errors.Is(sql.ErrNoRows, err) {
			return nil, model.UserErrNotFound
		}
		return nil, fmt.Errorf("failed find user: %w at mysql.UpdateUser", err)
	}

	usr.Name = []byte(name)
	usr.Password = password
	_, err = usr.Update(ctx, tx, boil.Infer())
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("failed update user: %w at mysql.UpdateUser", err)
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("failed update user: %w at mysql.UpdateUser", err)
	}
	return &model.User{
		ID:   uid,
		Name: name,
	}, nil
}
