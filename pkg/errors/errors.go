package errors

import (
	"fmt"
)

// #Dev errors

// ErrNotImplementedYet is raised when a resource is not implemented yet.
type ErrNotImplementedYet struct {
	Version string
}

func (err ErrNotImplementedYet) Error() string {
	return fmt.Sprintf("not implemented in version %s", err.Version)
}

// ErrNotFound is raised when a mandatory resource is not found in storage.
type ErrNotFound struct {
	Store string
	Index string
}

func (err ErrNotFound) Error() string {
	return fmt.Sprintf("no results found in store %s for index %s", err.Store, err.Index)
}

// ErrInvalidField is raised when a param is invalid.
type ErrInvalidField struct {
	Field string
	Value string
}

func (err ErrInvalidField) Error() string {
	return fmt.Sprintf("invalid field %s (%s)", err.Field, err.Value)
}
