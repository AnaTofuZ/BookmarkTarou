package model

// User ユーザー情報
type User struct {
	ID   uint64 `db:"id"`
	Name string `db:"name"`
}

// Users Userの集合(テーブル名と対応)
type Users []User
