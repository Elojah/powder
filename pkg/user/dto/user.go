package dto

import (
	"github.com/elojah/powder/pkg/errors"
	"github.com/elojah/powder/pkg/ulid"
)

// LoginReq request format for POST /login
type LoginReq struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	APIKey string `json:"api_key"`
}

// Check checks params validity.
func (req LoginReq) Check() error {
	if _, err := ulid.Parse(req.ID); err != nil {
		return errors.ErrInvalidField{Field: "id", Value: req.ID}
	}
	return nil
}
