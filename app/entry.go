package app

import (
	"context"
	"fmt"
	"github.com/anatofuz/BookmarkTarou/infra/store"
	"github.com/anatofuz/BookmarkTarou/model"
)

type BookmarkApp interface {
	Create(ctx context.Context, uid uint64, url, comment string) (*model.Entry, error)
	ListWithBCount(ctx context.Context) (*[]model.EntryWithBCount, error)
}

type BookmarkAppImpl struct {
	entryStore    store.EntryStore
	bookmarkStore store.BookmarkStore
}

func (b *BookmarkAppImpl) ListWithBCount(ctx context.Context) (*[]model.EntryWithBCount, error) {
	ewbs, err := b.entryStore.ListWirhBCount(ctx, 0, 20)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("failed ListWithBCount : %w", err)
	}
	return ewbs, nil
}

func NewBookmarkApp(entryStore store.EntryStore, bookmarkStore store.BookmarkStore) BookmarkApp {
	return &BookmarkAppImpl{entryStore: entryStore, bookmarkStore: bookmarkStore}
}

func (b *BookmarkAppImpl) Create(ctx context.Context, uid uint64, url, comment string) (*model.Entry, error) {
	etr, err := b.entryStore.FindEntryFromURL(ctx, url)
	if err != nil {
		return nil, fmt.Errorf("failed Create: %w at app.CreateBookmarkImpl", err)
	}
	if etr == nil {
		title, err := getTitleFromURL(url)
		if err != nil {
			title = url
		}
		etr, err = b.entryStore.Create(ctx, url, title)
		if err != nil {
			return nil, fmt.Errorf("failed Create: %w at app.CreateBookmarkImpl", err)
		}
	}
	_, err = b.bookmarkStore.Create(ctx, uid, etr.ID, comment)
	if err != nil {
		return nil, fmt.Errorf("failed Create Bookmark: %w at app.CreateBookmarkImpl", err)
	}
	return etr, nil
}
