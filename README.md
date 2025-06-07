# IM-Go

IM-Go 是一个基于 Go 语言开发的即时通讯系统。本系统提供高性能、可扩展的即时通讯解决方案，支持单聊、群聊等多种通讯场景。

## 项目架构

![](https://s2.loli.net/2025/05/26/7Nj1o58FP2axUY9.png)

## 功能特性

- **WebSocket 实时连接**：借助 WebSocket 技术，实现客户端与服务器之间的实时双向通信，确保消息即时送达，为用户带来流畅、无延迟的聊天体验。
- **单聊与群聊支持**：提供一对一的私密聊天和多人参与的群组聊天功能，满足用户在不同社交场景下的交流需求。
- **消息持久化存储**：将用户发送的消息持久化存储到数据库中，方便用户随时查看历史聊天记录，确保重要信息不丢失。
- **Kafka 消息队列**：引入 Kafka 作为消息队列，实现消息的异步处理和分发，提高系统的吞吐量和稳定性，确保在高并发场景下也能高效运行。
- **RESTful API 接口**：提供标准的 RESTful API 接口，方便第三方系统集成，拓展系统的应用范围和功能。
- **日志管理系统**：采用完善的日志管理系统，记录系统运行过程中的关键信息，便于开发人员进行问题排查和系统监控。
- **在线状态管理**：实时显示用户的在线状态，让用户随时了解好友或群成员的在线情况，提升沟通效率。
- **消息已读未读状态**：支持消息已读未读状态的标记，让用户清楚知晓对方是否查看了自己发送的消息，增强沟通的互动性。
- **消息历史记录查询**：用户可以方便地查询历史聊天记录，快速回顾之前的交流内容。
- **多种消息类型**：支持文本、图片等多种消息类型，满足用户多样化的表达需求。
- **用户关系管理**：提供用户关系管理功能，包括添加好友、管理好友列表等，方便用户构建自己的社交网络。
- **群组管理**：支持创建、管理群组，包括设置群主、群公告、群成员管理等，方便用户组织和管理群组活动。

## 技术栈

1. 编程语言
Go 1.24.2：采用高效、并发性能优越的 Go 语言进行开发，确保系统具备出色的性能和可扩展性。
2. Web 框架
Gin Web 框架：轻量级、高性能的 Gin Web 框架，用于构建 RESTful API 接口，提高开发效率。
3. 数据库操作
GORM：强大的 Go 语言 ORM 库，简化数据库操作，提高代码的可维护性。
4. 数据库存储
MySQL：使用 MySQL 作为关系型数据库，存储用户信息、消息记录等重要数据。
5. 实时通讯协议
WebSocket：基于 WebSocket 协议实现实时通讯，确保消息的即时传输。
6. 消息队列系统
Kafka：高性能的分布式消息队列 Kafka，用于处理和分发消息，提升系统的吞吐量和稳定性。
7. 数据序列化
proto buffer：采用 proto buffer 进行数据序列化，减少数据传输量，提高数据传输效率。
8. 日志记录
Zap 日志库：高性能的 Zap 日志库，用于记录系统运行日志，方便开发人员进行问题排查和系统监控。
9. 配置管理
Viper 配置管理：使用 Viper 进行配置管理，方便对系统的各种配置进行统一管理和维护。

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
```sql
-- 创建数据库
CREATE DATABASE chat;

-- 使用数据库
USE chat;

-- 导入 chat.sql 文件中的 SQL 语句
SOURCE /path/to/chat.sql;
```
- 修改配置文件：编辑 configs 目录下的配置文件 configs.toml，根据实际情况修改数据库连接信息和 Kafka 配置。
```toml
[mysql]
host = "127.0.0.1"
name = "chat"
password = "your_password"
port = 3306
tablePrefix = ""
user = "root"

[msgChannelType]
channelType = "gochannel"
kafkaHosts = "kafka:9092"
kafkaTopic = "go-chat-message"
```

4. 运行项目
```bash
go run ./cmd/main.go
```

## 配置说明

项目使用 Viper 进行配置管理，配置文件位于 `configs` 目录下。主要配置项包括：

- 数据库连接信息
```toml
[mysql]
host = "数据库主机地址"
name = "数据库名称"
password = "数据库密码"
port = 数据库端口号
tablePrefix = "表前缀"
user = "数据库用户名"
```

- Kafka配置
[msgChannelType]
channelType = "消息队列类型（gochannel 或 kafka）"
kafkaHosts = "Kafka 主机地址"
kafkaTopic = "Kafka 主题名称"

## 项目结构

```
.
├── api               # API 接口定义
│   ├── file_controller.go  # 文件相关接口
│   ├── group_controller.go # 群组相关接口
│   ├── message_controller.go # 消息相关接口
│   └── user_controller.go  # 用户相关接口
├── chat.sql          # 数据库结构定义
├── cmd               # 主程序入口
│   └── main.go
├── configs           # 配置文件
│   └── configs.toml
├── internal          # 内部包
│   ├── config        # 配置管理
│   │   └── config.go
│   ├── dao           # 数据访问层
│   │   └── pool
│   │       └── mysql_pool.go
│   ├── kafka         # Kafka 相关操作
│   │   ├── consumer.go
│   │   └── producer.go
│   ├── model         # 数据模型
│   │   ├── group.go
│   │   ├── group_member.go
│   │   ├── message.go
│   │   ├── user.go
│   │   └── user_friend.go
│   ├── router        # 路由管理
│   │   ├── router.go
│   │   └── socket.go
│   ├── server        # 服务器端逻辑
│   │   ├── client.go
│   │   └── server.go
│   └── service       # 业务逻辑层
│       ├── group_service.go
│       ├── message_service.go
│       └── user_service.go
└── pkg               # 公共包
    ├── common        # 通用工具和常量
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
    ├── errors        # 错误处理
    │   └── error.go
    ├── log           # 日志管理
    │   └── logger.go
    └── protocol      # 协议定义
        ├── message.pb.go
        └── message.proto
```