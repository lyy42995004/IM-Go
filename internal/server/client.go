package server

import (
	"github.com/gorilla/websocket"
	"github.com/lyy42995004/IM-Go/pkg/common/constant"
	"github.com/lyy42995004/IM-Go/pkg/log"
	"github.com/lyy42995004/IM-Go/pkg/protocol"
	"google.golang.org/protobuf/proto"
)

type Client struct {
	Conn *websocket.Conn
	Name string
	Send chan []byte // 向客户端发送消息
}

// 从客户端的 WebSocket 连接中读取消息，并根据消息类型进行相应的处理
func (c *Client) Write() {
	defer func() {
		MyServer.Ungister <- c // 注销登录
		c.Conn.Close()         // 关闭连接
	}()

	// 消息读取
	for {
		c.Conn.PongHandler()
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			log.Error("client read message error", log.Err(err))
			MyServer.Ungister <- c
			c.Conn.Close()
			break
		}

		msg := &protocol.Message{}
		proto.Unmarshal(message, msg)

		// 处理心跳消息
		if msg.Type == constant.HEAT_BEAT {
			pong := &protocol.Message{
				Content: constant.PONG,
				Type:    constant.HEAT_BEAT,
			}
			pongBytes, err2 := proto.Marshal(pong)
			if err2 != nil {
				log.Error("client marshal message error", log.Err(err))
			}
			c.Conn.WriteMessage(websocket.BinaryMessage, pongBytes)
		} else {
			MyServer.Broadcast <- message
		}
	}
}

// 从 Send 通道中接收消息，并将其以二进制消息的形式发送给客户端
func (c *Client) Read() {
	defer func() {
		c.Conn.Close()
	}()

	for message := range c.Send {
		c.Conn.WriteMessage(websocket.BinaryMessage, message)
	}
}
