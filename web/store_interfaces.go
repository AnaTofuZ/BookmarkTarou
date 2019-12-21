package web

import "github.com/anatofuz/BookmarkTarou/infra/store"

type StoreInterfaces struct {
	BookmarkStore    store.BookmarkStore
	EntryStore       store.EntryStore
	UserSessionStore store.UserSessionStore
	UserStore        store.UserStore
}
