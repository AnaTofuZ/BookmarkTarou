package mysql

import (
	"context"
	"database/sql"
	"github.com/anatofuz/BookmarkTarou/infra/store"
	"github.com/anatofuz/BookmarkTarou/model"
)

type entryStoreImpl struct {
	db *sql.DB
}

func NewEntryStore(db *sql.DB) store.EntryStore {
	return &entryStoreImpl{db: db}
}

func (e *entryStoreImpl) Create(ctx context.Context, url, title string) (*model.Entry, error) {
	panic("implement me")
}

func (e *entryStoreImpl) BookmarkedCountFromEntryIDs(ctx context.Context, ids *[]uint64) (*[]uint64, error) {
	panic("implement me")
}
