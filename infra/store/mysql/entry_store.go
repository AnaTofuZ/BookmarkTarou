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

func (e *entryStoreImpl) BookmarkedCountFromEntryIDs(ctx context.Context, ids *[]uint64) (*[]model.EntryWithBCount, error) {
	//ents, err := record.Entries(qm.Where("id=?",ids),qm.Load("bookmarks")).All(ctx,e.db)
	//if err != nil {
	//	return nil, fmt.Errorf("failed get entry... :%w",err)
	//}
	return nil, nil
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

func (e *entryStoreImpl) List(ctx context.Context, offset, limit int) (*[]model.Entry, error) {
	ets, err := record.Entries(qm.Offset(offset), qm.Limit(limit)).All(ctx, e.db)
	if err != nil {
		return nil, fmt.Errorf("failed get list entry: %w", err)
	}
	modelEntries := make([]model.Entry, 0, len(ets))
	for _, et := range ets {
		modelEntries = append(modelEntries, model.Entry{
			ID:    et.ID,
			Title: et.Titile,
			URL:   string(et.URL),
		})
	}
	return &modelEntries, nil
}

func (e *entryStoreImpl) ListWirhBCount(ctx context.Context, offset, limit int) (*[]model.EntryWithBCount, error) {
	ets, err := record.Entries(qm.Offset(offset), qm.Limit(limit), qm.Load("Bookmarks")).All(ctx, e.db)
	if err != nil {
		return nil, fmt.Errorf("failed get list entry: %w", err)
	}
	modelEntriesWC := make([]model.EntryWithBCount, 0, len(ets))
	for _, et := range ets {
		modelEntriesWC = append(modelEntriesWC, model.EntryWithBCount{
			Entry: &model.Entry{
				ID:    et.ID,
				Title: et.Titile,
				URL:   string(et.URL),
			},
			Count: uint64(len(et.R.Bookmarks)),
		})
	}
	return &modelEntriesWC, nil
}

func (e *entryStoreImpl) FindFromID(ctx context.Context, id uint64) (*model.Entry, error) {
	etr, err := record.FindEntry(ctx, e.db, id)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("failed FindFFromID: %w", err)
	}
	return &model.Entry{
		ID:    etr.ID,
		URL:   string(etr.URL),
		Title: etr.Titile,
	}, nil
}
