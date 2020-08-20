package dto

import "encoding/json"

type Method uint8

const (
	Login Method = iota
	Logout
	CreateRoom
	FetchRooms
	JoinRoom
	FetchRoomUsers
	Send
	SendUser
)

type M struct {
	Method  Method
	Content json.RawMessage
}
