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
	SignUp(ctx context.Context, name, password string) (*model.User, error)
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

func (u *UserAppImpl) SignUp(ctx context.Context, name, password string) (*model.User, error) {
	if name == "" {
		return nil, errors.New("empty user name")
	}
	if password == "" {
		return nil, errors.New("empty user password")
	}

	usrWP, err := u.userStore.GetPasswordWithUserFromName(ctx,name)
	if err != nil {
		return nil,fmt.Errorf("failed sign in: %w")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(password),usrWP.Pw); err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed signup: %w",err)
	}
	return &model.User{
		ID:   usrWP.ID,
		Name:usrWP.Name,
	}, nil
}
