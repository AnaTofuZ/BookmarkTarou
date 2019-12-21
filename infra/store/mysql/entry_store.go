package mysql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/anatofuz/BookmarkTarou/infra/record"
	"github.com/anatofuz/BookmarkTarou/infra/store"
	"github.com/anatofuz/BookmarkTarou/model"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

type entryStoreImpl struct {
	db *sql.DB
}

func NewEntryStore(db *sql.DB) store.EntryStore {
	return &entryStoreImpl{db: db}
}

func (e *entryStoreImpl) Create(ctx context.Context, url, title string) (*model.Entry, error) {
	et := &record.Entry{
		URL:    []byte(url),
		Titile: title,
	}
	tx, err := e.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed begin tx: %w at mysq/entry_store.Create", err)
	}
	if err := et.Insert(ctx, tx, boil.Infer()); err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("failed insert tx: %w at mysq/entry_store.Create", err)
	}
	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("failed commit tx: %w at mysq/entry_store.Create", err)
	}
	return &model.Entry{
		ID:    et.ID,
		URL:   string(et.URL),
		Title: et.Titile,
	}, nil
}

func (e *entryStoreImpl) BookmarkedCountFromEntryIDs(ctx context.Context, ids *[]uint64) (*[]uint64, error) {
	panic("implement me")
}

func (e *entryStoreImpl) FindEntryFromURL(ctx context.Context, url string) (*model.Entry, error) {
	et, err := record.Entries(qm.Where("url=?", url)).One(ctx, e.db)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed find entry: %w", err)
	}
	return &model.Entry{
		ID:    et.ID,
		URL:   string(et.URL),
		Title: et.Titile,
	}, nil
}
