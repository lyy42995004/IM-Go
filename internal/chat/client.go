package chat

import "github.com/gorilla/websocket"

type Client struct {
	Conn *websocket.Conn
	Name string
	Send chan []byte
}