package store

import (
	"context"
	"github.com/anatofuz/BookmarkTarou/model"
)

type BookmarkStore interface {
	Create(ctx context.Context, uid, entryID uint64, comment string) (*model.Bookmark, error)
	ListFromUID(ctx context.Context, uid, offset, limit uint64) (*[]model.Bookmark, error)
}
