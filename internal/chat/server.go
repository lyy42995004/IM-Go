package chat

import "sync"

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
	
}
