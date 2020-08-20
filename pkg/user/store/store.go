package store

import (
	"github.com/elojah/powder/pkg/user"
	"github.com/elojah/redis"
)

var _ user.Store = (*Store)(nil)
var _ user.StoreName = (*Store)(nil)

// Store implements user store.
type Store struct {
	*redis.Service
}

// NewStore returns a new user store.
func NewStore(s *redis.Service) *Store {
	return &Store{
		Service: s,
	}
}
