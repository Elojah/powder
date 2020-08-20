package store

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-redis/redis"

	gerrors "github.com/elojah/powder/pkg/errors"
	"github.com/elojah/powder/pkg/message"
)

const (
	messageKey = "message:"
)

// Upsert implementation for message in redis.
func (s *Store) Upsert(ctx context.Context, m message.M) error {
	raw, err := m.Marshal()
	if err != nil {
		return err
	}

	if err := s.Set(messageKey+m.RoomID.String()+":"+m.ID.String(), raw, 0).Err(); err != nil {
		return fmt.Errorf("upsert message %s: %w", m.ID.String(), err)
	}

	return nil
}

// Fetch implementation for message in redis.
func (s *Store) Fetch(ctx context.Context, filter message.Filter) (message.M, error) {
	val, err := s.Get(messageKey + filter.RoomID.String() + ":" + filter.ID.String()).Result()
	if err != nil {
		if !errors.Is(err, redis.Nil) {
			return message.M{}, fmt.Errorf("fetch message %s: %w", filter.ID.String(), err)
		}

		return message.M{}, fmt.Errorf(
			"fetch message %s: %w",
			filter.ID.String(),
			gerrors.ErrNotFound{Store: messageKey, Index: filter.ID.String() + ":" + filter.ID.String()},
		)
	}

	var m message.M
	if err := m.Unmarshal([]byte(val)); err != nil {
		return message.M{}, fmt.Errorf("fetch message %s: %w", filter.ID.String(), err)
	}

	return m, nil
}

// FetchMany implementation for message in redis.
func (s *Store) FetchMany(ctx context.Context, filter message.Filter) (map[string]message.M, error) {
	keys, err := s.Keys(messageKey + filter.RoomID.String() + ":*").Result()
	if err != nil {
		if !errors.Is(err, redis.Nil) {
			return nil, fmt.Errorf("fetch messages %s: %w", filter.RoomID.String(), err)
		}

		return nil, fmt.Errorf(
			"fetch messages %s: %w",
			filter.RoomID.String(),
			gerrors.ErrNotFound{Store: messageKey, Index: filter.RoomID.String() + ":*"},
		)
	}

	ms := make(map[string]message.M, len(keys))

	for _, key := range keys {
		val, err := s.Get(key).Result()
		if err != nil {
			return nil, fmt.Errorf("fetch message %s: %w", key, err)
		}

		var tmp message.M
		if err := tmp.Unmarshal([]byte(val)); err != nil {
			return nil, fmt.Errorf("fetch message %s: %w", key, err)
		}

		ms[tmp.ID.String()] = tmp
	}

	return ms, nil
}

// Remove implementation for message in redis.
func (s *Store) Remove(ctx context.Context, filter message.Filter) error {
	if err := s.Del(messageKey + filter.RoomID.String() + ":" + filter.ID.String()).Err(); err != nil {
		return fmt.Errorf("remove message %s: %w", filter.ID.String(), err)
	}

	return nil
}
