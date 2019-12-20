package store

import (
	"context"

	"github.com/anatofuz/BookmarkTarou/model"
)

// UserStore Userに関するDBアクセスのInterface
type UserStore interface {
	GetUsers(ctx context.Context) (*model.Users, error)
	GetUserFromID(ctx context.Context, uid uint64) (*model.User, error)
	CreateUser(ctx context.Context, name string, password string) (*model.User, error)
	UpdateUser(ctx context.Context, uid uint64, name string, password string) (*model.User, error)
}
