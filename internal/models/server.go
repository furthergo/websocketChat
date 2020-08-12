package models

import (
	"fmt"
	logger "github.com/futhergo/websocketChat/log"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"io"
	"net"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type wsChatServer struct {
	lis net.Listener
	upgrader *websocket.Upgrader
	clients []*wsChatClient
	addr string
	lock *sync.Mutex
}

func NewWsChatServer() *wsChatServer {
	return &wsChatServer{
		upgrader: &websocket.Upgrader{
			ReadBufferSize: 1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		addr: ":80",
		lock: &sync.Mutex{},
	}
}

func (s *wsChatServer)MessageAPI(c *gin.Context) {
	conn, err := s.upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		logger.Log2file("upgrade failed")
		return
	}
	client, err := s.Accept(c, conn)
	if err != nil {
		logger.Log2file("parse user id failed")
		return
	}
	s.BroadCast(1, fmt.Sprintf("New User coming in: %v", client.user.Name))
	s.lock.Lock()
	s.clients = append(s.clients, client)
	s.lock.Unlock()
	go s.Serve(client)
}

func (s *wsChatServer)Accept(c *gin.Context, conn *websocket.Conn) (*wsChatClient, error) {
	id, err := strconv.ParseUint(c.Query("uid"), 10, 64)
	if err != nil {
		return nil, err
	}
	u, err := GetUserById(uint(id))
	if err != nil {
		return nil, err
	}
	u.Ip = conn.RemoteAddr().String()
	return &wsChatClient{
		conn: conn,
		user:  u,
	}, nil
}

func (s *wsChatServer)Disconnect(c *wsChatClient) {
	s.BroadCast(1, fmt.Sprintf("[%v] quit!!!!!!!!!!!!!", c.user.Name))
	s.lock.Lock()
	for i, cc := range s.clients {
		if cc == c {
			s.clients = append(s.clients[:i], s.clients[i+1:]...)
			break
		}
	}
	s.lock.Unlock()
}

func (s *wsChatServer)Serve(c *wsChatClient) {
	go func() {
		historyMsgs := AllMessages()
		for _, m := range historyMsgs {
			c.WriteMessage(m)
		}
	}()

	for {
		t, b, err := c.conn.ReadMessage()
		if err == io.EOF {
			s.Disconnect(c)
			c.conn.Close()
			logger.Log2file( fmt.Sprintf("%v connect close", c.user.Name))
			break
		}
		if b == nil {
			s.Disconnect(c)
			c.conn.Close()
			break
		}
		if err != nil {
			s.Disconnect(c)
			c.conn.Close()
			logger.Log2file(err)
		}
		msg := string(b)
		m := Message{
			FromUid: c.user.ID,
			FromName: c.user.Name,
			Content: msg,
			SendTime: time.Now(),
			SendStatus: 1,
		}
		m.Save()
		go s.BroadCast(t, m.FormatString())
	}
}

func (s *wsChatServer)BroadCast(t int, msg string)  {
	logger.Log2file(msg)
	for _, c := range s.clients {
		c.Write(t, []byte(msg))
	}
}