package main

import log "github.com/lyy42995004/IM-Go/pkg/zap"

func main() {
	// 不需要显式初始化，第一次使用会自动初始化
	defer log.Sync()

	// 直接使用包级函数记录日志
	log.Info("Application started")
	log.Debug("Debug information") // 由于默认级别是info，这条不会输出
	
	// 带上下文的日志
	log.Error("Failed to process request",
		log.String("request_id", "abc123"),
		log.Int("attempt", 3))	
}
