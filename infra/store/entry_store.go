package store

import (
	"context"
	"github.com/anatofuz/BookmarkTarou/model"
)

type EntryStore interface {
	FindFromID(ctx context.Context, id uint64) (*model.Entry, error)
	Create(ctx context.Context, url, title string) (*model.Entry, error)
	List(ctx context.Context, offset, limit int) (*[]model.Entry, error)
	ListWirhBCount(ctx context.Context, offset, limit int) (*[]model.EntryWithBCount, error)
	BookmarkedCountFromEntryIDs(ctx context.Context, ids *[]uint64) (*[]model.EntryWithBCount, error)
	FindEntryFromURL(ctx context.Context, url string) (*model.Entry, error)
}
