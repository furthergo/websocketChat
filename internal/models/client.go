package models

import (
	"github.com/gorilla/websocket"
	"time"
)

type wsChatClient struct {
	conn *websocket.Conn
	user UserEntity
}

func (c *wsChatClient)Write(t int, b []byte) {
	c.conn.WriteMessage(t, b)
}

func getFormatTime() string {
	return time.Now().Format("15:04:05")
}