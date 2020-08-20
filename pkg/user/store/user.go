package store

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-redis/redis"

	gerrors "github.com/elojah/powder/pkg/errors"
	"github.com/elojah/powder/pkg/user"
)

const (
	userKey = "user:"
)

// Upsert implementation for user in redis.
func (s *Store) Upsert(ctx context.Context, u user.U) error {
	raw, err := u.Marshal()
	if err != nil {
		return err
	}

	if err := s.Set(userKey+u.ID.String(), raw, 0).Err(); err != nil {
		return fmt.Errorf("upsert user %s: %w", u.ID.String(), err)
	}

	return nil
}

// Fetch implementation for user in redis.
func (s *Store) Fetch(ctx context.Context, filter user.Filter) (user.U, error) {
	val, err := s.Get(userKey + filter.ID.String()).Result()
	if err != nil {
		if !errors.Is(err, redis.Nil) {
			return user.U{}, fmt.Errorf("fetch user %s: %w", filter.ID.String(), err)
		}

		return user.U{}, fmt.Errorf(
			"fetch user %s: %w",
			filter.ID.String(),
			gerrors.ErrNotFound{Store: userKey, Index: filter.ID.String() + ":" + filter.ID.String()},
		)
	}

	var u user.U
	if err := u.Unmarshal([]byte(val)); err != nil {
		return user.U{}, fmt.Errorf("fetch user %s: %w", filter.ID.String(), err)
	}

	return u, nil
}

// FetchMany implementation for user in redis.
func (s *Store) FetchMany(ctx context.Context, filter user.Filter) (map[string]user.U, error) {
	us := make(map[string]user.U, len(filter.IDs))

	for _, id := range filter.IDs {
		val, err := s.Get(userKey + id.String()).Result()
		if err != nil {
			return nil, fmt.Errorf("fetch user %s: %w", id.String(), err)
		}

		var tmp user.U
		if err := tmp.Unmarshal([]byte(val)); err != nil {
			return nil, fmt.Errorf("fetch users %s: %w", filter.ID.String(), err)
		}

		us[tmp.ID.String()] = tmp
	}

	return us, nil
}

// Remove implementation for user in redis.
func (s *Store) Remove(ctx context.Context, filter user.Filter) error {
	if err := s.Del(userKey + filter.ID.String()).Err(); err != nil {
		return fmt.Errorf("remove user %s: %w", filter.ID.String(), err)
	}

	return nil
}
