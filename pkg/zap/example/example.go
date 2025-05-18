package main

import (
	"errors"

	log "github.com/lyy42995004/IM-Go/pkg/zap"
)

func main() {
	// 1. 基础日志
	log.Info("服务启动")
	log.Warn("内存使用量较高")

	// 2. 带参数的日志
	log.Info("用户登录",
		log.String("用户名", "张三"),
		log.Int("登录次数", 3),
	)

	// 3. 错误日志
	err := errors.New("连接超时")
	log.Error("请求失败", log.Err(err))

	// 4. 同步日志（确保所有日志都写入磁盘）
	log.Sync()
}