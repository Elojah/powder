package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"

	perrors "github.com/elojah/powder/pkg/errors"
	"github.com/elojah/powder/pkg/ulid"
	"github.com/elojah/powder/pkg/user"
	userdto "github.com/elojah/powder/pkg/user/dto"
	"github.com/rs/zerolog/log"
	"nhooyr.io/websocket"
)

func (h handler) login(ctx context.Context, c *websocket.Conn, msg json.RawMessage) {

	logger := log.With().Str("method", "login").Logger()

	// #Request processing
	var request userdto.LoginReq
	if err := json.NewDecoder(bytes.NewBuffer(msg)).Decode(&request); err != nil {
		logger.Error().Err(err).Msg("invalid payload")
		return
	}

	if err := request.Check(); err != nil {
		logger.Error().Err(err).Msg("invalid payload")
		return
	}

	userID := ulid.MustParse(request.ID)

	// #Check duplicate user
	// TO IMPROVE
	if _, err := h.user.Fetch(ctx, user.Filter{ID: userID}); err != nil {
		if !errors.As(err, &perrors.ErrNotFound{}) {
			logger.Error().Err(err).Msg("user already connected")
			return
		}
	}
	if _, err := h.user.FetchName(ctx, request.Name); err != nil {
		if !errors.As(err, &perrors.ErrNotFound{}) {
			logger.Error().Err(err).Msg("name already exists")
			return
		}
	}

	// #Create new user
	if err := h.user.Upsert(ctx, user.U{
		ID:   userID,
		Name: request.Name,
	}); err != nil {
		logger.Error().Err(err).Msg("failed to upsert user")
		return
	}

	_ = c.Write(ctx, websocket.MessageText, []byte("successfully login"))
}
