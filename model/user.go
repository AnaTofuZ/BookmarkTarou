package model

import "errors"

// User ユーザー情報
type User struct {
	ID   uint64
	Name string
}

// UserWithPW パスワード付きのモデル
type UserWithPW struct {
	User
	Pw []byte
}

// Users Userの集合(テーブル名と対応)
type Users []User

// UserErrNotFound ユーザーが見つからない場合のエラー
var UserErrNotFound = errors.New("user:not found")
