package main

import (
	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"
	"nhooyr.io/websocket"

	"github.com/elojah/powder/pkg/message/dto"
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

		log.Info().Msg(fmt.Sprintf("received: %s", string(buf)))

		var msg dto.M
		if err := msg.Unmarshal(buf); err != nil {
			logger.Error().Err(err).Msg("failed to read message")
			http.Error(w, fmt.Sprintf("failed to read message: %v", err), http.StatusInternalServerError)

			return
		}

		switch msg.Method {
		case dto.Logout:
			h.logout(ctx, c, msg.Content)
		case dto.CreateRoom:
			h.createRoom(ctx, c, msg.Content)
		case dto.FetchRooms:
			h.fetchRooms(ctx, c, msg.Content)
		case dto.JoinRoom:
			h.joinRoom(ctx, c, msg.Content)
		case dto.FetchRoomUsers:
			h.fetchRoomUsers(ctx, c, msg.Content)
		case dto.Send:
			h.send(ctx, c, msg.Content)
		case dto.SendUser:
			h.sendUser(ctx, c, msg.Content)

		}
	}
}
