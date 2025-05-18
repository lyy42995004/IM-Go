package zap

import (
	"fmt"
	"io"
	"strings"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"gopkg.in/natefinch/lumberjack.v2"
)

type RotateConfig struct {
	// 共用配置
	Filename string // 完整文件名
	MaxAge   int    // 保留旧日志文件的最大天数

	// 按时间轮转配置
	RotationTime time.Duration // 日志文件轮转时间

	// 按大小轮转配置
	MaxSize    int  // 日志文件最大大小（MB）
	MaxBackups int  // 保留日志文件的最大数量
	Compress   bool // 是否对日志文件进行压缩归档
	LocalTime  bool // 是否使用本地时间，默认 UTC 时间
}

// NewProductionRotateByTime 创建按时间轮转的 io.Writer
func NewProductionRotateByTime(filename string) io.Writer {
	return NewRotateByTime(NewProductionRotateConfig(filename))
}

// NewProductionRotateBySize 创建按大小轮转的 io.Writer
func NewProductionRotateBySize(filename string) io.Writer {
	return NewRotateBySize(NewProductionRotateConfig(filename))
}

func NewProductionRotateConfig(filename string) *RotateConfig {
	return &RotateConfig{
		Filename: filename,
		MaxAge:   30, // 日志保留 30 天

		RotationTime: time.Hour * 24, // 24 小时轮转一次

		MaxSize:    100, // 100M
		MaxBackups: 100, // 多于100个日志文件后，清理较旧的日志
		Compress:   true,
		LocalTime:  true,
	}
}

func NewRotateByTime(cfg *RotateConfig) io.Writer {
	opts := []rotatelogs.Option{
		rotatelogs.WithMaxAge(time.Duration(cfg.MaxAge) * 24 * time.Hour),
		rotatelogs.WithRotationTime(cfg.RotationTime), // 关键参数
		rotatelogs.WithLinkName(cfg.Filename),         // 保持软链接指向最新
	}

	// 时间格式处理逻辑
	filenameParts := strings.SplitN(cfg.Filename, ".", 2)
	pattern := fmt.Sprintf("%s.%%Y-%%m-%%d-%%H-%%M-%%S.%s",
		filenameParts[0], filenameParts[1])

	writer, _ := rotatelogs.New(pattern, opts...)
	return writer
}

func NewRotateBySize(cfg *RotateConfig) io.Writer {
	return &lumberjack.Logger{
		Filename:   cfg.Filename,
		MaxSize:    cfg.MaxSize,    // 触发轮转的阈值(MB)
		MaxBackups: cfg.MaxBackups, // 经济型存储策略
		Compress:   cfg.Compress,   // 节省磁盘空间
		LocalTime:  cfg.LocalTime,
	}
}
