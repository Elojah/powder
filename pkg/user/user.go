package user

import (
	"context"

	"github.com/elojah/powder/pkg/ulid"
)

// Filter object to fetch specific U.
type Filter struct {
	ID  ulid.ID
	IDs []ulid.ID
}

// App application layer for U object.
type App interface {
	Store
}

// Store storage layer for U object.
type Store interface {
	Upsert(context.Context, U) error
	Fetch(context.Context, Filter) (U, error)
	FetchMany(context.Context, Filter) (map[string]U, error)
	Remove(context.Context, Filter) error
}
