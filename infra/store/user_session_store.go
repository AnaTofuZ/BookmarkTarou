package store

import "github.com/anatofuz/BookmarkTarou/model"

type UserSessionStore interface {
	Get(token string) (*model.User, error)
	Set(token string, usrSession *model.User) error
}
