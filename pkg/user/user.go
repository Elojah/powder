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
	StoreName
}

// Store storage layer for U object.
type Store interface {
	Upsert(context.Context, U) error
	Fetch(context.Context, Filter) (U, error)
	FetchMany(context.Context, Filter) (map[string]U, error)
	Remove(context.Context, Filter) error
}

// Store storage layer for user name index.
type StoreName interface {
	UpsertName(context.Context, string) error
	FetchName(context.Context, string) (bool, error)
}
