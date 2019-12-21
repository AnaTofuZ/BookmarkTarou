package app

import (
	"context"
	"errors"
	"fmt"

	"github.com/anatofuz/BookmarkTarou/infra/store"
	"github.com/anatofuz/BookmarkTarou/model"
	"golang.org/x/crypto/bcrypt"
)

type UserApp interface {
	Create(ctx context.Context, name, password string) (*model.User, error)
	SignIn(ctx context.Context, name, password string) (*model.User, error)
}

type UserAppImpl struct {
	userStore store.UserStore
}

func NewUserApp(userStore store.UserStore) UserApp {
	return &UserAppImpl{userStore: userStore}
}

func (u *UserAppImpl) Create(ctx context.Context, name, password string) (*model.User, error) {
	if name == "" {
		return nil, errors.New("empty user name")
	}
	if password == "" {
		return nil, errors.New("empty user password")
	}
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed CreateUser: %w", err)
	}

	usr, err := u.userStore.CreateUser(ctx, name, hashedPass)
	if err != nil {
		return nil, fmt.Errorf("failed createuser: %w", err)
	}
	return usr, nil
}

func (u *UserAppImpl) SignIn(ctx context.Context, name, password string) (*model.User, error) {
	if name == "" {
		return nil, errors.New("empty user name")
	}
	if password == "" {
		return nil, errors.New("empty user password")
	}
	usrWP, err := u.userStore.GetPasswordWithUserFromName(ctx, name)
	fmt.Print(usrWP)
	if err != nil {
		return nil, fmt.Errorf("faile get pw: %w", err)
	}

	if err := bcrypt.CompareHashAndPassword(usrWP.Pw,[]byte(password)); err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return nil, nil
		}
		return nil, fmt.Errorf("failed login: %w", err)
	}
	return &usrWP.User, nil
}
