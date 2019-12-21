package mysql

import (
	"context"
	"database/sql"
	"github.com/anatofuz/BookmarkTarou/infra/store"
	"github.com/anatofuz/BookmarkTarou/model"
)

type bookmarkStoreImpl struct {
	db *sql.DB
}

func NewBookmarkStore(db *sql.DB) store.BookmarkStore {
	return &bookmarkStoreImpl{db: db}
}

func (b *bookmarkStoreImpl) Create(ctx context.Context, uid, entryID uint64, comment string) (*model.Bookmark, error) {
	panic("implement me")
}

func (b *bookmarkStoreImpl) ListFromUID(ctx context.Context, uid, offset, limit uint64) (*[]model.Bookmark, error) {
	panic("implement me")
}
