package server

import (
	"encoding/base64"
	"os"
	"strings"
	"sync"

	"github.com/google/uuid"
	"github.com/lyy42995004/IM-Go/internal/config"
	"github.com/lyy42995004/IM-Go/internal/service"
	"github.com/lyy42995004/IM-Go/pkg/common/constant"
	"github.com/lyy42995004/IM-Go/pkg/common/util"
	"github.com/lyy42995004/IM-Go/pkg/log"
	"github.com/lyy42995004/IM-Go/pkg/protocol"
	"google.golang.org/protobuf/proto"
)

var MyServer = NewServer()

// 服务器实例
type Server struct {
	Clients   map[string]*Client // 存储已连接的客户端信息
	mutex     *sync.Mutex
	Broadcast chan []byte  // 广播消息
	Register  chan *Client // 处理客户端注册请求
	Ungister  chan *Client // 处理客户端注销请求
}

// 创建服务器实例
func NewServer() *Server {
	return &Server{
		Clients:   make(map[string]*Client),
		mutex:     &sync.Mutex{},
		Broadcast: make(chan []byte),
		Register:  make(chan *Client),
		Ungister:  make(chan *Client),
	}
}

func (s *Server) Start() {
	log.Info("start server...")
	for {
		select {
		// 处理客户端注册
		case conn := <-s.Register:
			log.Info("login", log.String("login", conn.Name))
			s.Clients[conn.Name] = conn
			msg := &protocol.Message{
				From:    "System",
				To:      conn.Name,
				Content: "welcome!",
			}
			protoMsg, _ := proto.Marshal(msg) // 序列化为字节切片
			conn.Send <- protoMsg

		// 处理客户端注销
		case conn := <-s.Ungister:
			log.Info("loginout", log.String("loginout", conn.Name))
			if _, ok := s.Clients[conn.Name]; ok {
				close(conn.Send)
				delete(s.Clients, conn.Name)
			}

		// 处理消息广播
		case message := <-s.Broadcast:
			msg := &protocol.Message{}
			proto.Unmarshal(message, msg)

			if msg.To != "" {
				// 有指定接收者的消息处理
				if msg.ContentType >= constant.TEXT && msg.ContentType <= constant.VIDEO {
					// 一般消息，比如文本消息，视频文件消息等
					_, exists := s.Clients[msg.From]
					if exists { // 检查发送者是否在连接列表中
						saveMessage(msg)
					}

					if msg.ContentType == constant.MESSAGE_TYPE_USER {
						// 单人消息处理
						s.sendUserMessage(msg)
					} else if msg.ContentType == constant.MESSAGE_TYPE_GROUP {
						// 多人消息处理
						s.sendGroupMessage(msg)
					} else {
						// 语音电话，视频电话等，仅支持单人聊天，不支持群聊
						// 不保存文件，直接进行转发
						client, ok := s.Clients[msg.To]
						if ok {
							client.Send <- message
						}
					}
				} else {
					// 无指定接收者的广播消息处理
					for id, conn := range s.Clients {
						log.Info("allUser", log.String("allUser", id))

						select {
						case conn.Send <- message: // 发送消息给客户端，成功继续处理
						default: // 失败关闭客户端
							close(conn.Send)
							delete(s.Clients, conn.Name)
						}
					}
				}
			}
		}
	}
}

// 发送字节切片并发送给接收者
func (s *Server) sendUserMessage(msg *protocol.Message) {
	client, ok := s.Clients[msg.To]
	if ok {
		msgBytes, err := proto.Marshal(msg)
		if err == nil {
			client.Send <- msgBytes
		}
	}
}

// 将消息发送给群组中的所有成员
func (s *Server) sendGroupMessage(msg *protocol.Message) {
	// 获取所有成员信息
	users := service.GroupService.GetUserIdByGroupUuid(msg.To)
	for _, user := range users {
		if user.Uuid == msg.From {
			continue
		}

		client, ok := s.Clients[user.Uuid]
		if !ok {
			continue
		}
		
		fromUserDetails := service.UserService.GetUserDetails(msg.From)
		msgSend := protocol.Message{
			Avatar:       fromUserDetails.Avatar,
			FromUsername: msg.FromUsername,
			From:         msg.To,
			To:           msg.From,
			Content:      msg.Content,
			ContentType:  msg.ContentType,
			Type:         msg.Type,
			MessageType:  msg.MessageType,
			Url:          msg.Url,
		}

		msgByte, err := proto.Marshal(&msgSend)
		if err == nil {
			client.Send <- msgByte
		}
	}
}

// 存消息，如果是文本消息直接保存，如果是文件，语音等消息，保存文件后，保存对应的文件路径
func saveMessage(message *protocol.Message) {
	// 如果上传的是base64字符串文件，解析文件保存
	if message.ContentType == 2 {
		url := uuid.New().String() + ".png"
		index := strings.Index(message.Content, "base64")
		index += 7

		content := message.Content
		content = content[index:]

		dataBuffer, dataErr := base64.StdEncoding.DecodeString(content)
		if dataErr != nil {
			log.Error("transfer base64 to file error", log.String("transfer base64 to file error", dataErr.Error()))
			return
		}
		err := os.WriteFile(config.GetConfig().StaticPath.FilePath+url, dataBuffer, 0666)
		if err != nil {
			log.Error("write file error", log.String("write file error", err.Error()))
			return
		}
		message.Url = url
		message.Content = ""
	} else if message.ContentType == 3 {
		// 普通的文件二进制上传
		fileSuffix := util.GetFileType(message.File)
		nullStr := ""
		if nullStr == fileSuffix {
			fileSuffix = strings.ToLower(message.FileSuffix)
		}
		contentType := util.GetContentTypeBySuffix(fileSuffix)
		url := uuid.New().String() + "." + fileSuffix
		err := os.WriteFile(config.GetConfig().StaticPath.FilePath+url, message.File, 0666)
		if err != nil {
			log.Error("write file error", log.String("write file error", err.Error()))
			return
		}
		message.Url = url
		message.File = nil
		message.ContentType = contentType
	}

	service.MessageService.SaveMessage(message)
}

// 消费kafka里面的消息, 直接放入go channel中统一进行消费
func ConsumerKafkaMsg(data []byte) {
	MyServer.Broadcast <- data
}
