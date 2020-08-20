package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"
	"nhooyr.io/websocket"

	messagedto "github.com/elojah/powder/pkg/message/dto"
)

func (h handler) connect(w http.ResponseWriter, r *http.Request) {

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
		if websocket.CloseStatus(err) == websocket.StatusNormalClosure {
			return
		}

		var buf []byte
		_, buf, err = c.Read(ctx)

		if err != nil {
			logger.Error().Err(err).Msg("failed to read websocket")
			http.Error(w, fmt.Sprintf("failed to read websocket: %v", err), http.StatusInternalServerError)

			return
		}

		// log.Info().Msg(fmt.Sprintf("received: %s", string(buf)))

		var msg messagedto.M
		//  if err := msg.Unmarshal(buf); err != nil {
		if err := json.NewDecoder(bytes.NewBuffer(buf)).Decode(&msg); err != nil {
			logger.Error().Err(err).Msg("failed to read message")
			http.Error(w, fmt.Sprintf("failed to read message: %v", err), http.StatusInternalServerError)

			return
		}

		switch msg.Method {
		case messagedto.Login:
			h.login(ctx, c, msg.Content)
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
