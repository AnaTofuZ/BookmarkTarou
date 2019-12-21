package store

import (
	"context"

	"github.com/anatofuz/BookmarkTarou/model"
)

// UserStore Userに関するDBアクセスのInterface
type UserStore interface {
	GetUserFromID(ctx context.Context, uid uint64) (*model.User, error)
	GetPasswordWithUserFromName(ctx context.Context, name string) (*model.UserWithPW, error)
	CreateUser(ctx context.Context, name string, password []byte) (*model.User, error)
	UpdateUser(ctx context.Context, uid uint64, name string, password []byte) (*model.User, error)
}
