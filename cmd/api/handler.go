package main

import (
	"context"
	"net/http"

	"github.com/elojah/powder/pkg/message"
	"github.com/elojah/powder/pkg/user"
	"github.com/rs/zerolog/log"
)

type handler struct {
	srv *http.Server

	apikey string

	user    user.App
	message message.App
}

// Dial starts the auth server.
func (h *handler) Dial(c Config) error {
	h.apikey = c.APIKey

	mux := http.NewServeMux()

	mux.HandleFunc("/login", h.Login)
	// mux.HandleFunc("/logout", h.Logout)
	// mux.HandleFunc("/room/create", h.CreateRoom)
	// mux.HandleFunc("/rooms", h.FetchRooms)
	// mux.HandleFunc("/room/join", h.CreateRoom)
	// mux.HandleFunc("/room/users", h.FetchRoomUsers)
	// mux.HandleFunc("/send", h.Send)
	// mux.HandleFunc("/user/send", h.Send)

	h.srv = &http.Server{
		Addr:    c.Address,
		Handler: mux,
	}
	go func() {
		if err := h.srv.ListenAndServeTLS(c.Cert, c.Key); err != nil {
			log.Error().Err(err).Msg("failed to start server")
		}
	}()
	return nil
}

// Close shutdowns the server listening.
func (h *handler) Close() error {
	return h.srv.Shutdown(context.Background())
}
