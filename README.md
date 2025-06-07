# IM-Go

IM-Go 是一个基于 Go 语言开发的即时通讯系统。本系统提供高性能、可扩展的即时通讯解决方案，支持单聊、群聊等多种通讯场景。

## 功能特性

- 基于WebSocket的实时通讯
- 支持单聊和群聊
- 消息持久化存储
- 使用Kafka进行消息队列
- RESTful API接口
- 日志管理系统
- 在线状态管理
- 消息已读未读状态
- 消息历史记录查询
- 支持文本、图片等多种消息类型
- 用户关系管理
- 群组管理功能

## 技术栈

- Go 1.24.2
- Gin Web框架
- GORM
- MySQL
- WebSocket
- Kafka
- proto buffer
- Zap日志库
- Viper配置管理

## 快速开始

### 前置要求

- Go 1.24.2 或更高版本
- MySQL
- Kafka

### 安装

1. 克隆项目
```bash
git clone https://github.com/lyy42995004/IM-Go.git
cd IM-Go
```

2. 安装依赖
```bash
go mod download
```

3. 配置数据库
- 创建数据库并导入 `chat.sql`
- 修改 `configs` 目录下的配置文件

4. 运行项目
```bash
go run ./cmd/main.go
```

## 配置说明

项目使用 Viper 进行配置管理，配置文件位于 `configs` 目录下。主要配置项包括：

- 数据库连接信息
- Kafka配置

## 项目结构

```
.
├── api       # API接口定义
│   ├── file_controller.go
│   ├── group_controller.go
│   ├── message_controller.go
│   └── user_controller.go
└── chat.sql  # 数据库结构
├── cmd       # 主程序入口
│   └── main.go
├── configs   # 配置文件
│   └── configs.toml
├── internal  # 内部包
│   ├── config
│   │   └── config.go
│   ├── dao
│   │   └── pool
│   │       └── mysql_pool.go
│   ├── kafka
│   │   ├── consumer.go
│   │   └── producer.go
│   ├── model
│   │   ├── group.go
│   │   ├── group_member.go
│   │   ├── message.go
│   │   ├── user.go
│   │   └── user_friend.go
│   ├── router
│   │   ├── router.go
│   │   └── socket.go
│   ├── server
│   │   ├── client.go
│   │   └── server.go
│   └── service
│       ├── group_service.go
│       ├── message_service.go
│       └── user_service.go
└── pkg       # 公共包
    ├── common
    │   ├── constant
    │   │   └── constant.go
    │   ├── request
    │   │   ├── friend_request.go
    │   │   └── message_request.go
    │   ├── response
    │   │   ├── group_response.go
    │   │   ├── message_response.go
    │   │   ├── response_msg.go
    │   │   └── search_response.go
    │   └── util
    │       └── util.go
    ├── errors
    │   └── error.go
    ├── log
    │   └── logger.go
    └── protocol
        ├── message.pb.go
        └── message.proto
```