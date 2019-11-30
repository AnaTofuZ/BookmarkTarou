package model

// User ユーザー情報
type User struct {
	ID   uint64 `db:"id"`
	Name string `db:"name"`
}
