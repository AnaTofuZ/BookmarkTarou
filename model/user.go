package model

import "errors"

// User ユーザー情報
type User struct {
	ID   uint64
	Name string
}

// Users Userの集合(テーブル名と対応)
type Users []User

// UserErrNotFound ユーザーが見つからない場合のエラー
var UserErrNotFound = errors.New("user:not found")
