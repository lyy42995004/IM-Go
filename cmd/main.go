package main

import (
	log "github.com/lyy42995004/IM-Go/pkg/zap"
)

func main() {
	log.Info("日志test",
		log.String("用户名", "张三"),
		log.Int("登录次数", 3))
	log.Sync()
}
