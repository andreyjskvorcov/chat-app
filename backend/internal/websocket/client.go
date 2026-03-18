package websocket

import "github.com/gorilla/websocket"

type Client struct {
	conn   *websocket.Conn
	userID string
	room   string
}
