package zap_test

import (
	"errors"
	"os"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/lyy42995004/IM-Go/pkg/zap"
)

func TestAllLogLevels(t *testing.T) {
	tests := []struct {
		name    string
		logFunc func(string, ...zap.Field)
		level   string
	}{
		{"Debug", zap.Debug, "debug"},
		{"Info", zap.Info, "info"},
		{"Warn", zap.Warn, "warn"},
		{"Error", zap.Error, "error"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.logFunc("测试"+tt.name+"级别日志", 
				zap.String("test_case", tt.name),
				zap.Time("timestamp", time.Now()))
			
			// 可以在这里添加断言，检查日志文件是否包含预期内容
		})
	}
}

func TestLogFileRotation(t *testing.T) {
	// 备份原始文件
	origPath := zap.GetLogPath()
	defer func() {
		zap.SetLogPath(origPath)
	}()

	// 使用临时目录测试
	tmpDir := t.TempDir()
	testPath := tmpDir + "/test.log"
	zap.SetLogPath(testPath)

	// 写入足够多的日志以触发轮转
	for i := 0; i < 5000; i++ {
		zap.Info("日志轮转测试",
			zap.Int("count", i),
			zap.String("data", strings.Repeat("a", 2000))) // 每条约2KB
	}

	// 检查日志文件是否创建
	if _, err := os.Stat(testPath); os.IsNotExist(err) {
		t.Errorf("日志文件未创建: %v", err)
	}

	// 检查是否有备份文件创建（根据实际轮转配置）
	// ...
}

func TestConcurrentSafety(t *testing.T) {
	var wg sync.WaitGroup
	count := 100

	wg.Add(count)
	for i := 0; i < count; i++ {
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				zap.Info("并发安全测试",
					zap.Int("goroutine", id),
					zap.Int("count", j))
			}
		}(i)
	}
	wg.Wait()

	// 检查日志文件是否有损坏或丢失的条目
	// ...
}

// 注意：需要在实际代码中暴露SetLogPath方法用于测试
func (z *zapLogger) SetLogPath(path string) {
	logPath = path
	// 需要重新初始化logger
	init()
}