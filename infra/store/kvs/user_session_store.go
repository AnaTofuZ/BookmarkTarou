package kvs

import (
	"encoding/json"
	"fmt"
	"github.com/anatofuz/BookmarkTarou/infra/store"
	"github.com/anatofuz/BookmarkTarou/model"
	"github.com/go-redis/redis"
	"time"
)

type redisStore struct {
	client *redis.Client
}


func NewRedisStore(client *redis.Client) store.UserSessionStore{
	return &redisStore{client:client}
}

func (r redisStore) Get(token string) (*model.User, error) {
	var user model.User

	busr, err := r.client.Get(token).Bytes()
	if err != nil {
		return nil, fmt.Errorf("failed Get session: %w",err)
	}
	if err := json.Unmarshal(busr,&user); err != nil {
		return nil, fmt.Errorf("failed Get session: %w",err)
	}
	return &user,nil
}

func (r redisStore) Set(token string, usrSession *model.User) error {
	busr, err := json.Marshal(usrSession)
	if err != nil {
		return fmt.Errorf("failed to save session: %w",err)
	}

	if err := r.client.Set(token, busr, 24*time.Hour).Err(); err != nil {
		return fmt.Errorf("failed to save session: %w",err)
	}
	return nil
}
