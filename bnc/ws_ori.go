package bnc

import "github.com/gorilla/websocket"

type Ws struct {
	conn *websocket.Conn
}
