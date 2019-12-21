package store

import (
	"context"
	"github.com/anatofuz/BookmarkTarou/model"
)

type EntryStore interface {
	Create(ctx context.Context, url, title string) (*model.Entry, error)
	BookmarkedCountFromEntryIDs(ctx context.Context, ids *[]uint64) (*[]uint64, error)
}
