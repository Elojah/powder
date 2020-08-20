package store

import (
	"github.com/elojah/powder/pkg/message"
	"github.com/elojah/redis"
)

var _ message.Store = (*Store)(nil)

// Store implements message store.
type Store struct {
	*redis.Service
}

// NewStore returns a new message store.
func NewStore(s *redis.Service) *Store {
	return &Store{
		Service: s,
	}
}
