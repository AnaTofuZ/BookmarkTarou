package model

// Bookmark ユーザーがブックマークしたもの。コメントがある
type Bookmark struct {
	ID      uint64 `db:"id"`
	Comment string `db:"comment"`
	UserID  uint64 `db:"user_id"`
	EntryID uint64 `db:"entry_id"`
}
