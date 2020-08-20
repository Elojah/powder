package message

import (
	"context"

	"github.com/elojah/powder/pkg/ulid"
)

// Filter object to fetch specific M.
type Filter struct {
	ID     ulid.ID
	RoomID ulid.ID
}

// App application layer for M object.
type App interface {
	Store
}

// Store storage layer for M object.
type Store interface {
	Upsert(context.Context, M) error
	Fetch(context.Context, Filter) (M, error)
	FetchMany(context.Context, Filter) (map[string]M, error)
	Remove(context.Context, Filter) error
}
