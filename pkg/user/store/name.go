package store

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-redis/redis"
)

const (
	usernameKey = "nameuser:"
)

// UpsertName implementation for username in redis.
func (s *Store) UpsertName(ctx context.Context, name string) error {

	if err := s.HSet(usernameKey, name, true).Err(); err != nil {
		return fmt.Errorf("upsert username %s: %w", name, err)
	}

	return nil
}

// FetchName implementation for user in redis.
func (s *Store) FetchName(ctx context.Context, name string) (bool, error) {
	_, err := s.Get(usernameKey + name).Result()
	if err != nil {
		if !errors.Is(err, redis.Nil) {
			return false, fmt.Errorf("fetch username %s: %w", name, err)
		}

		return false, nil
	}

	return true, nil
}
