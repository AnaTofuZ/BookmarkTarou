package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/anatofuz/BookmarkTarou/infra/record"
	"github.com/anatofuz/BookmarkTarou/infra/store"
	"github.com/anatofuz/BookmarkTarou/model"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

type bookmarkStoreImpl struct {
	db *sql.DB
}

func NewBookmarkStore(db *sql.DB) store.BookmarkStore {
	return &bookmarkStoreImpl{db: db}
}

func (b *bookmarkStoreImpl) Create(ctx context.Context, uid, entryID uint64, comment string) (*model.Bookmark, error) {
	bk := &record.Bookmark{
		UserID:  uid,
		EntryID: entryID,
		Comment: comment,
	}
	tx, err := b.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed begin tx: %w at mysq/bookmark_store.Create", err)
	}
	if err := bk.Insert(ctx, tx, boil.Infer()); err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("failed begin tx: %w at mysq/bookmark_store.Create", err)
	}
	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("failed begin tx: %w at mysq/bookmark_store.Create", err)
	}
	return &model.Bookmark{
		ID:      bk.ID,
		Comment: bk.Comment,
		UserID:  bk.UserID,
		EntryID: bk.EntryID,
	}, nil
}

func (b *bookmarkStoreImpl) ListFromUID(ctx context.Context, uid, offset, limit uint64) (*[]model.Bookmark, error) {
	panic("implement me")
}

func (b *bookmarkStoreImpl) ListFromEntryID(ctx context.Context, entryID uint64) (*[]model.Bookmark, error) {
	bookmarks, err := record.Bookmarks(qm.Where("entry_id = ?", entryID)).All(ctx, b.db)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("failed: %w", err)
	}

	modelBookmarks := make([]model.Bookmark, 0, len(bookmarks))

	for _, bkm := range bookmarks {
		modelBookmarks = append(modelBookmarks, model.Bookmark{
			ID:      bkm.ID,
			Comment: bkm.Comment,
			UserID:  bkm.UserID,
			EntryID: bkm.EntryID,
		})
	}
	return &modelBookmarks, nil
}

func (b *bookmarkStoreImpl) ListFromEntryIDWUsers(ctx context.Context, entryID uint64) (*[]model.Bookmark, *[]model.User, error) {
	bookmarks, err := record.Bookmarks(qm.Where("entry_id = ?", entryID), qm.Load("User")).All(ctx, b.db)
	if err != nil {
		fmt.Println(err)
		return nil, nil, fmt.Errorf("failed: %w", err)
	}
	modelBookmarks := make([]model.Bookmark, 0, len(bookmarks))
	modelUsers := make([]model.User, 0, len(bookmarks))

	for _, bkm := range bookmarks {
		modelBookmarks = append(modelBookmarks, model.Bookmark{
			ID:      bkm.ID,
			Comment: bkm.Comment,
			UserID:  bkm.UserID,
			EntryID: bkm.EntryID,
		})
		modelUsers = append(modelUsers, model.User{
			ID:   bkm.R.User.ID,
			Name: string(bkm.R.User.Name),
		})
	}
	return &modelBookmarks, &modelUsers, nil
}
