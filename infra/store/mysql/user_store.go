package mysql

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/anatofuz/BookmarkTarou/infra/store"
	"github.com/anatofuz/BookmarkTarou/model"
)

type userStoreImpl struct {
	db *sqlx.DB
}

func NewUserStore(db *sqlx.DB) store.UserStore {
	return &userStoreImpl{db: db}
}

func (u userStoreImpl) GetUsers(ctx context.Context) (*model.Users, error) {
	panic("implement me")
}

func (u userStoreImpl) GetUserFromID(ctx context.Context, uid uint64) (*model.User, error) {
	panic("implement me")
}

func (u userStoreImpl) CreateUser(ctx context.Context, name string, password string) (*model.User, error) {
	panic("implement me")
}

func (u userStoreImpl) UpdateUser(ctx context.Context, uid uint64, name string, password string) (*model.User, error) {
	panic("implement me")
}

