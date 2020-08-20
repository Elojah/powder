package app

import "github.com/elojah/powder/pkg/user"

// App user logic implementation.
type App struct {
	user.Store
	user.StoreName
}
