package models

import (
	"fmt"
	logger "github.com/futhergo/websocketChat/log"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"io"
	"net"
	"net/http"
	"strings"
	"sync"
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
	}
	client := s.Accept(c, conn)
	s.BroadCast(1, fmt.Sprintf("New User coming in: %v", client.user.Name))
	s.lock.Lock()
	s.clients = append(s.clients, client)
	s.lock.Unlock()
	go s.Serve(client)
}

func (s *wsChatServer)Accept(c *gin.Context, conn *websocket.Conn) *wsChatClient {
	n := c.Query("username")
	p := c.Query("password")
	var u UserEntity
	u.Name = n
	u.Password = p
	u.Ip = c.Request.RemoteAddr
	return &wsChatClient{
		conn: conn,
		user:  u,
	}
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
		if strings.HasPrefix(msg, "NAME:") {
			msg = strings.TrimPrefix(msg, "NAME:")
			logger.Log2file(fmt.Sprintf("[%v] set new name: [%v]", c.user.Name, msg))
			c.user.Name = msg
		} else {
			msg = fmt.Sprintf("[%v](%v): %v", c.user.Name, getFormatTime(), string(b))
			go s.BroadCast(t, msg)
		}
	}
}

func (s *wsChatServer)BroadCast(t int, msg string)  {
	logger.Log2file(msg)
	for _, c := range s.clients {
		c.Write(t, []byte(msg))
	}
}