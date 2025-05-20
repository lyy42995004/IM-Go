package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/lyy42995004/IM-Go/internal/server"
	"github.com/lyy42995004/IM-Go/pkg/log"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// 建立WebSocket连接，并开启客户端的读写操作
func RunSocket(c *gin.Context) {
	user := c.Query("user")
	if user == "" {
		return
	}
	log.Info("newUser", log.String("newUser", user))

	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}

	client := &server.Client{
		Name: user,
		Conn: ws,
		Send: make(chan []byte),
	}

	server.MyServer.Register <- client
	go client.Read()
	go client.Write()
}
