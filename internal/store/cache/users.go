package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/marcosmcb/backend-engineering-with-go/internal/store"
	"github.com/redis/go-redis/v9"
)

type UserStore struct {
	rdb *redis.Client
}

const UserExpireTime = time.Minute * 15

func (s *UserStore) Get(ctx context.Context, userID int64) (*store.User, error) {
	cacheKey := fmt.Sprintf("userid-%d", userID)
	data, err := s.rdb.Get(ctx, cacheKey).Result()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	var user store.User
	if data != "" {
		err := json.Unmarshal([]byte(data), &user)
		if err != nil {
			return nil, err
		}
	}
	return &user, nil
}

func (s *UserStore) Set(ctx context.Context, user *store.User) error {
	cacheKey := fmt.Sprintf("userid-%d", user.ID)
	val, err := json.Marshal(user)

	if err != nil {
		return err
	}
	return s.rdb.SetEx(ctx, cacheKey, val, UserExpireTime).Err()
}
