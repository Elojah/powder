package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"
	"nhooyr.io/websocket"

	perrors "github.com/elojah/powder/pkg/errors"
	messagedto "github.com/elojah/powder/pkg/message/dto"
	"github.com/elojah/powder/pkg/ulid"
	"github.com/elojah/powder/pkg/user"
	userdto "github.com/elojah/powder/pkg/user/dto"
)

func (h handler) Login(w http.ResponseWriter, r *http.Request) {

	// TODO this should be done on handler side, not inside controller
	switch r.Method {
	case http.MethodPost:
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	ctx := r.Context()
	logger := log.With().Str("route", r.URL.EscapedPath()).Str("method", r.Method).Str("address", r.RemoteAddr).Logger()

	// #Request processing
	var request userdto.LoginReq
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		logger.Error().Err(err).Msg("invalid payload")
		http.Error(w, fmt.Sprintf("invalid payload: %v", err), http.StatusBadRequest)

		return
	}

	if err := request.Check(); err != nil {
		logger.Error().Err(err).Msg("invalid payload")
		http.Error(w, fmt.Sprintf("invalid payload: %v", err), http.StatusBadRequest)

		return
	}

	userID := ulid.MustParse(request.ID)

	// #Check duplicate user
	// TO IMPROVE
	if _, err := h.user.Fetch(ctx, user.Filter{ID: userID}); err != nil {
		if !errors.As(err, &perrors.ErrNotFound{}) {
			logger.Error().Err(err).Msg("user already connected")
			http.Error(w, fmt.Sprintf("user already connected: %v", err), http.StatusBadRequest)

			return
		}
	}
	if _, err := h.user.FetchName(ctx, request.Name); err != nil {
		if !errors.As(err, &perrors.ErrNotFound{}) {
			logger.Error().Err(err).Msg("name already exists")
			http.Error(w, fmt.Sprintf("name already exists: %v", err), http.StatusBadRequest)

			return
		}
	}

	// #Create new user
	if err := h.user.Upsert(ctx, user.U{
		ID:   userID,
		Name: request.Name,
	}); err != nil {
		logger.Error().Err(err).Msg("failed to upsert user")
		http.Error(w, fmt.Sprintf("failed to upsert user: %v", err), http.StatusBadRequest)

		return

	}

	c, err := websocket.Accept(w, r, nil)
	if err != nil {
		logger.Error().Err(err).Msg("failed to accept websocket")
		http.Error(w, fmt.Sprintf("failed to accept websocket: %v", err), http.StatusInternalServerError)

		return
	}
	defer c.Close(websocket.StatusInternalError, "failed to close websocket")

	for {
		var buf []byte
		_, buf, err = c.Read(ctx)

		if err != nil {
			logger.Error().Err(err).Msg("failed to read websocket")
			http.Error(w, fmt.Sprintf("failed to read websocket: %v", err), http.StatusInternalServerError)

			return
		}

		// log.Info().Msg(fmt.Sprintf("received: %s", string(buf)))

		var msg messagedto.M
		if err := msg.Unmarshal(buf); err != nil {
			logger.Error().Err(err).Msg("failed to read message")
			http.Error(w, fmt.Sprintf("failed to read message: %v", err), http.StatusInternalServerError)

			return
		}

		switch msg.Method {
		case messagedto.Logout:
			h.logout(ctx, c, msg.Content)
		case messagedto.CreateRoom:
			h.createRoom(ctx, c, msg.Content)
		case messagedto.FetchRooms:
			h.fetchRooms(ctx, c, msg.Content)
		case messagedto.JoinRoom:
			h.joinRoom(ctx, c, msg.Content)
		case messagedto.FetchRoomUsers:
			h.fetchRoomUsers(ctx, c, msg.Content)
		case messagedto.Send:
			h.send(ctx, c, msg.Content)
		case messagedto.SendUser:
			h.sendUser(ctx, c, msg.Content)
		default:
			_ = c.Write(ctx, websocket.MessageText, []byte("unknown method"))
		}
	}
}
